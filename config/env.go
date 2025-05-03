package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBDomain   string
	Port       string
}

var Env EnvVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(".env missing")
	}

	Env = EnvVariables{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDomain:   os.Getenv("DB_DOMAIN"),
		Port:       os.Getenv("PORT"),
	}
}
