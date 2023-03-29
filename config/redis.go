package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisConfiguration struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func RdbConfiguration() *redis.Options {
	host := viper.GetString("REDIS_HOST")
	port := viper.GetString("REDIS_PORT")
	password := viper.GetString("REDIS_PASSWORD")
	db := viper.GetInt("REDIS_DB")

	addr := fmt.Sprintf("%s:%s", host, port)

	config := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}

	return config
}
