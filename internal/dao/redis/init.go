package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func Init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_, err := RedisCli.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}