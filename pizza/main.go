package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/VarthanV/pizza/handlers"
	"github.com/VarthanV/pizza/users"
	"github.com/VarthanV/pizza/users/implementation"
	"github.com/VarthanV/pizza/users/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

func  main(){
	var db *sql.DB
	{
		var err error
		// Initialize mysql database
		//<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
		db, err = sql.Open("mysql","root:root123@tcp(localhost:3306)/pizza_shop")
		if err != nil {
			glog.Error("Unable to connect to db...",err)
			os.Exit(-1)
		}	
	}
	glog.Info("Connected to mysql db.....")
	var redisClient *redis.Client
	{
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		// Make a ping
		ping := redisClient.Ping(context.TODO())
		result , _ := ping.Result()
		glog.Info("Result from redis ping...",result)
	}
	var usersvc users.Service
	{
		dbRepository, err := repositories.NewMySqlRepository(db)
		rdbRepository := repositories.NewRedisRepository(redisClient)
		if err != nil {
			glog.Error("Unable to initialize dbRepository...",err)
			return
		}
		if rdbRepository !=nil {
			glog.Error("Unable to initialize rdb Repository....")
		}
		glog.Info("Initializing user service....")
		usersvc = implementation.NewService(dbRepository,rdbRepository)
	}

	glog.Info("Init handlers....")
	userHandler := handlers.NewUserHandler(usersvc)
	router := gin.Default()
	//TODO: Mapping signup route with its respective handler.... Will be refactored later
	router.POST("/signup",userHandler.SignUpUserHandler)
	router.GET("/test",userHandler.Test)
	router.POST("/login",userHandler.LoginUserHandler)

	// Run the router
	router.Run(":8080")


}