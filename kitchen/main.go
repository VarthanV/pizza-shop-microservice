package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/VarthanV/kitchen/cooks"
	implementation "github.com/VarthanV/kitchen/implementation"
	"github.com/VarthanV/kitchen/inmemorydb"
	"github.com/VarthanV/kitchen/migrations"
	"github.com/VarthanV/kitchen/mysql"
	"github.com/VarthanV/kitchen/processes"
	"github.com/VarthanV/kitchen/queue"
	"github.com/VarthanV/kitchen/redisclient"
	"github.com/VarthanV/kitchen/seeder"
	"github.com/VarthanV/kitchen/shared"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	var db *sql.DB
	{
		var err error
		err = godotenv.Load()
		if err != nil {
			glog.Fatalf("Unable to load environment variables")
		}
		// Initializing DB Constants
		dbConnection := shared.DBConnection{
			DBName:   os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
		}
		// Initialize mysql database
		//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			dbConnection.Username,
			dbConnection.Password,
			dbConnection.Host,
			dbConnection.Port,
			dbConnection.DBName)
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			glog.Fatalf("Unable to connect to db...", err)
			os.Exit(-1)
		}
	}
	glog.Info("Connected to mysql db.....")
	// Msg queue

	rabbitMqConnection, err := amqp.Dial(os.Getenv("RABBIT_MQ_CONNECTION_STRING"))
	if err != nil {
		glog.Fatalf("Unable to connect to rabbit mq %f", err)
	}

	// RMQ Channel
	ch, err := rabbitMqConnection.Channel()
	if err != nil {
		glog.Fatalf("Unable to create a channel %f", err)
	}

	migrationsvc := migrations.NewMigrationService(db)
	migrationsvc.RunMigrations(context.TODO())

	isSeedingEnabled := os.Getenv("SEEDING_ENABLED")
	// Seed data -> If not needed we can omit it env and write logic accordingly
	if isSeedingEnabled == "true" {
		seederSvc := seeder.NewSeederService(db)
		seederSvc.SeedData()
	}
	var redisClient *redis.Client
	{
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf(`%s:%s`, os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		// Make a ping
		ping := redisClient.Ping(context.TODO())
		result, err := ping.Result()
		if err != nil {
			glog.Info("Error connecting to redis...", err)
		}
		glog.Info("Result from redis ping...", result)
	}

	var cookservice cooks.Service
	{
		cookRepo := mysql.NewCookMysqlRepo(db)
		cookservice = implementation.NewCookService(cookRepo)
	}

	var processUpdateService processes.OrderProcessUpdateService
	{
		repo := mysql.NewOrderProcessUpdateRepoMysql(db)
		processUpdateService = implementation.NewOrderOrderProcessUpdateImplementation(repo)
	}
	var processOrderSvc processes.OrderProcessService
	{
		processOrderSvc = implementation.NewProcessOrderImplementationService(cookservice, processUpdateService)
	}
	var orderRequestInmemoryService inmemorydb.OrderRequestInMemoryService
	{
		repo := redisclient.NewOrderQueueRepo(redisClient)
		orderRequestInmemoryService = implementation.NewOrderInmemoryService(repo)
	}

	var orderRequestsvc processes.OrderRequestService
	{
		orderRequestsvc = implementation.NewOrderRequestImplementation(cookservice, processOrderSvc, orderRequestInmemoryService)
	}
	var queueService queue.QueueService
	{
		queueRepo := queue.NewRabbitRepository(ch)
		queueService = implementation.NewRabbitMQService(queueRepo, orderRequestsvc)
	}

	glog.Info(cookservice)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	ctx := context.Background()
	queueService.ConsumeOrderDetails(ctx)
	r.Run() // l
}
