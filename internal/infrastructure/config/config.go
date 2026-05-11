package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBUrl   string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using system env")
	}

	dbUrl := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASS") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") + "?sslmode=disable"

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   dbUrl,
	}
}
