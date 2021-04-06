package config

import (
	"belajariah-main-service/model"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnectDB to get all needed db connections for application
func ConnectDB(config *model.Config) *sqlx.DB {
	return getDBConnection(config)
}

func getDBConnection(config *model.Config) *sqlx.DB {
	dbConnectionStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=verify-full sslcert=ca-certificate.crt",
		config.Database.Host,
		config.Database.Port,
		config.Database.DbName,
		config.Database.Username,
		config.Database.Password,
	)

	db, err := sqlx.Open("postgres", dbConnectionStr)
	if err != nil {
		log.Panicln("Error establishing connection to *belajariah main* database", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("Error connecting to *port main* database", err)
		panic(err)
	}

	//TODO: experiment with correct values
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	return db
}
