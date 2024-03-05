package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-redis/redis/v9"
)

var RedisConn *redis.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	rd := os.Getenv("RedisDB")
	redis, err := GetRedisConnection(rd)
	if err != nil {
		fmt.Println("ERR IS", err)
	}
	RedisConn = redis

}

func GetRedisConnection(rd string) (*redis.Client, error) {

	opt, err := redis.ParseURL(rd)
	client := redis.NewClient(opt)
	defer client.Conn().Close()
	if err != nil {
		return nil, err
	}
	return client, nil
}
