package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DbConfig DbConfig
}

type DbConfig struct {
	Dsn string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("не удалось найти .env файл")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":8080"
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = "postgresql://admin:supersecret@localhost:5432/postgres"
	}

	return Config{
		Port: port,
		DbConfig: DbConfig{
			Dsn: dsn,
		},
	}
}
