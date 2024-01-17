package main

import (
	"be-service-public-api/config"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	_DeliveryHTTP "be-service-public-api/public-api/delivery/http"
	_RepoGRPCPublicAPI "be-service-public-api/public-api/repository/grpc"

	_RepoMySQLPublicAPI "be-service-public-api/public-api/repository/mysql"
	_UsecasePublicAPI "be-service-public-api/public-api/usecase"

	grpcpool "github.com/processout/grpc-go-pool"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	dbFlag = "mysql"
)

func main() {
	// CLI options parse
	configFile := flag.String("c", "config.yaml", "Config file")
	flag.Parse()

	// Config file
	config.ReadConfig(*configFile)

	// Set log level
	switch viper.GetString("server.log_level") {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	// Initialize database connection
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.database"))
	val := url.Values{}
	val.Add("multiStatements", "true")
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Migrate database if any new schema
	driver, err := mysql.WithInstance(dbConn, &mysql.Config{})
	if err == nil {
		mig, err := migrate.NewWithDatabaseInstance(viper.GetString("database.path_migrate"), viper.GetString("mysql.database"), driver)
		log.Info(viper.GetString("database.path_migrate"))
		if err == nil {
			err = mig.Up()
			if err != nil {
				if err == migrate.ErrNoChange {
					log.Debug("No database migration")
				} else {
					log.Error(err)
				}
			} else {
				log.Info("Migrate database success")
			}
			version, dirty, err := mig.Version()
			if err != nil && err != migrate.ErrNilVersion {
				log.Error(err)
			}
			log.Debug("Current DB version: " + strconv.FormatUint(uint64(version), 10) + "; Dirty: " + strconv.FormatBool(dirty))
		} else {
			log.Warn(err)
		}
	} else {
		log.Warn(err)
	}

	// Initialize Redis
	ctx := context.Background()
	dbRedis := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Username: viper.GetString("redis.username"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.database"),
		PoolSize: viper.GetInt("redis.max_connection"),
	})

	_, err = dbRedis.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Redis connection established")
	var grpcPoolProduct *grpcpool.Pool
	productConn := func() (client *grpc.ClientConn, err error) {
		address := fmt.Sprintf("%s:%s", viper.GetString("grpc.product_service.host"), viper.GetString("grpc.product_service.port"))
		client, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	grpcPoolProduct, err = grpcpool.New(productConn, viper.GetInt("grpc.init"), viper.GetInt("grpc.capacity"), time.Duration(viper.GetInt("grpc.idle_duration"))*time.Second, time.Duration(viper.GetInt("grpc.max_life_duration"))*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	// Register GRPC
	// repoGRPCAuth := _RepoGRPCAuth.NewGRPCAuthRepository(grpcPoolAuth)
	repoGRPCProduct := _RepoGRPCPublicAPI.NewGRPCProductRepository(grpcPoolProduct)
	// Register repository & usecase public API

	repoMySQLPublicAPI := _RepoMySQLPublicAPI.NewMySQLPublicAPIRepository(dbConn)
	usecasePublicAPI := _UsecasePublicAPI.NewPublicAPIUsecase(repoMySQLPublicAPI, repoGRPCProduct)
	// serverAuth := _RepoGRPCAuthObject.NewGRPCAuth(usecaseAuth)
	// Initialize gRPC server
	go func() {
		listen, err := net.Listen("tcp", ":"+viper.GetString("server.grpc_port"))
		if err != nil {
			log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
		}

		grpcServer := grpc.NewServer()
		// _RepoGRPCAuthServer.RegisterAuthorizationServiceServer(grpcServer, serverAuth)
		log.Println("gRPC server is running in port", viper.GetString("server.grpc_port"))
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Initialize HTTP web framework
	app := fiber.New(fiber.Config{
		Prefork:       viper.GetBool("server.prefork"),
		StrictRouting: viper.GetBool("server.strict_routing"),
		CaseSensitive: viper.GetBool("server.case_sensitive"),
		BodyLimit:     viper.GetInt("server.body_limit"),
	})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: viper.GetString("middleware.allows_origin"),
	}))

	// HTTP routing
	app.Get(viper.GetString("server.base_path")+"/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	_DeliveryHTTP.RouterAPI(app, usecasePublicAPI)

	// Start Fiber HTTP server
	if err := app.Listen(":" + viper.GetString("server.port")); err != nil {
		log.Fatal(err)
	}
}