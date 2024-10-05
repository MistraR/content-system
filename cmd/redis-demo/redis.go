package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := connectRDB()
	ctx := context.Background()
	err := rdb.Set(ctx, "session-id:admin", "session-id", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}
	result, err := rdb.Get(ctx, "session-id:admin").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		panic(err)
	}
	fmt.Println("查询session-id = ", result)
}

func connectRDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.132.113.82:6378",
		Password: "pP6vY4sD", // no password set
		DB:       0,          // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
