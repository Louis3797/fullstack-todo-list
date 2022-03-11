package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const TodoPrefix string = "todo:"

var Ctx = context.Background()

var RedisClient *redis.Client

func InitRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		DB:           0, // use default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected succesfully to Redis on Port: 6379")
}
