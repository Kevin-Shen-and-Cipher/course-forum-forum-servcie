package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	err error
)

// RdbConnection create redis connection
func RdbConnection(options *redis.Options) error {
	rdb = redis.NewClient(options)

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Redis connection error")
		return err
	}

	rdb.FlushAll(context.Background())

	return nil
}

// GetRedis connection
func GetRdb() *redis.Client {
	return rdb
}
