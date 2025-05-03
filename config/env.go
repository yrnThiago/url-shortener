package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	ClientUrl  string
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
		ClientUrl:  os.Getenv("CLIENT_URL"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBDomain:   os.Getenv("DB_DOMAIN"),
		Port:       os.Getenv("PORT"),
	}
}
