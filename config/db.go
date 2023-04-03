package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DBConfiguration() string {
	dbName := viper.GetString("db_NAME")
	dbUser := viper.GetString("db_USER")
	dbPassword := viper.GetString("db_PASSWORD")
	dbHost := viper.GetString("db_HOST")
	dbPort := viper.GetString("db_PORT")
	dbSslMode := viper.GetString("SSL_MODE")

	dbDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode,
	)

	return dbDSN
}
