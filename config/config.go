package config

import (
	"belajariah-main-service/model"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() *model.Config {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	mailPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	logMaxAge, _ := strconv.Atoi(os.Getenv("LOG_MAXAGE"))
	logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAXSIZE"))
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	logMaxBackup, _ := strconv.Atoi(os.Getenv("LOG_MAXBACKUP"))
	logCompress, _ := strconv.ParseBool(os.Getenv("LOG_COMPRESS"))

	config := &model.Config{
		Server: model.ServerConfig{
			Port: serverPort,
		},
		Database: model.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			DbName:   os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Log: model.LogConfig{
			Filename:   os.Getenv("LOG_FILENAME"),
			MaxSize:    logMaxSize,
			MaxBackups: logMaxBackup,
			MaxAge:     logMaxAge,
			Compress:   logCompress,
		},
		Mail: model.MailConfig{
			AuthEmail:    os.Getenv("AUTH_EMAIL"),
			AuthPassword: os.Getenv("AUTH_PASSWORD"),
			SMTPHost:     os.Getenv("SMTP_HOST"),
			SMTPPort:     mailPort,
			SenderName:   os.Getenv("SENDER_NAME"),
		},
	}

	return config
}
