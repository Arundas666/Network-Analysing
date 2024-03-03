package config

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

var RedisConn *redis.Client

func init() {
	redis, err := GetRedisConnection()
	if err != nil {
		fmt.Println(err)
	}
	RedisConn = redis
}

func GetRedisConnection() (*redis.Client, error) {
	opt, err := redis.ParseURL("rediss://default:ba29eda96217425197e2d53fc7ef30f6@us1-more-skylark-40395.upstash.io:40395")
	client := redis.NewClient(opt)
	// log.Println("Redis connection created")

	// ctx := context.Background()
	// timeOutCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	// defer cancel()
	// _, err := singletonRedis.Conn.Ping(timeOutCtx).Result()
	// if err != nil {
	// 	fmt.Printf("Error connecting to redis: %v\n", err)
	// 	singletonRedis = nil
	// 	return nil
	// }
	if err != nil {
		return nil, err
	}
	return client, nil
}
