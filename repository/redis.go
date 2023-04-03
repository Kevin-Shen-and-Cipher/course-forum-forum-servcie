package repository

import (
	"course-forum/infra/logger"
	rdb "course-forum/infra/redis"
	"time"

	redis "github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

var expiration = time.Hour * 1

func RedisSet(ctx *gin.Context, key string, val []byte) (err error) {
	_, err = rdb.GetRdb().Set(ctx, key, val, expiration).Result()

	if err != nil {
		logger.Fatalf("redis set error: %s \n", err.Error())
	} else {
		logger.Infof("redis set success\n")
	}

	return
}

func RedisGet(ctx *gin.Context, key string) (val string, err error) {
	val, err = rdb.GetRdb().Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			logger.Infof("redis not hit\n")
			return
		}
	}

	logger.Infof("redis hit!\n")

	return
}

func RedisDelete(ctx *gin.Context, key string) {
	_, err := rdb.GetRdb().Del(ctx, key).Result()

	if err != nil {
		logger.Fatalf("redis delete key error: %s \n", err.Error())
	} else {
		logger.Infof("redis key deleted!\n")
	}
}
