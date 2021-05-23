package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/VarthanV/kitchen/cooks"
	cookimpl "github.com/VarthanV/kitchen/cooks/implementation"
	"github.com/VarthanV/kitchen/mysql"
	"github.com/VarthanV/kitchen/queue"
	rimpl "github.com/VarthanV/kitchen/queue/implementation"
	"github.com/VarthanV/kitchen/shared"
	"github.com/gin-gonic/gin"
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

	var queueService queue.QueueService
	{
		queueRepo := queue.NewRabbitRepository(ch)
		queueService = rimpl.NewRabbitMQService(queueRepo)
	}

	var cookservice cooks.Service
	{
		cookRepo := mysql.NewCookMysqlRepo(db)
		cookservice = cookimpl.NewCookService(cookRepo)
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
