package database

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit() (err error) {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       1,
	})

	_, err = RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
	println("redis connect success !!!")
	return nil
}
