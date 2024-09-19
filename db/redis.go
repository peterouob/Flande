package db

import (
	"context"
	"ecomm/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.GetString("redis.addr"),
		Password: config.Config.GetString("redis.pass"),
		DB:       config.Config.GetInt("redis.db"),
	})

	ping := Rdb.Ping(context.Background())
	if ping.String() != "" {
		log.Println("Redis connection ...")
	}
}
