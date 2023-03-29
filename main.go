package main

import (
	"course-forum/config"
	"course-forum/infra/database"
	"course-forum/infra/logger"
	"course-forum/infra/redis"
	"course-forum/migrations"
	"course-forum/routers"
	"time"

	"github.com/spf13/viper"
)

// @title Course forum API
// @version 1.0
// @description This is the course forum api documentation

// @license.name MIT License
// @license.url https://github.com/Kevin-Shen-and-Cipher/course-forum-forum-servcie/blob/main/LICENSE
func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	dbDSN := config.DBConfiguration()

	if err := database.DbConnection(dbDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	options := config.RdbConfiguration()

	if err := redis.RdbConnection(options); err != nil {
		logger.Fatalf("redis RdbConnection error: %s", err)
	}

	//later separate migration
	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
