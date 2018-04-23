package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

type DBConfig struct {
	Dialect string
	Host    string
	DBName  string
	DBUser  string
	DBPass  string
	SSLMode string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Dialect: "postgres",
		Host:    os.Getenv("DB_HOST"),
		DBName:  os.Getenv("DB_NAME"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		SSLMode: os.Getenv("DB_SSL_MODE"),
	}
}

func NewDB(config *DBConfig) *Database {

	source := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.DBUser, config.DBPass, config.DBName)

	db, err := sqlx.Connect(config.Dialect, source)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{DB: db}
}
