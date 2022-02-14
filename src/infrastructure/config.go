package infrastructure

import (
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	DBName string
	Host string
	User string
	Password string
	Port int
}

func NewDBConfig() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return &DBConfig{
		DBName: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port: port,
	}
}