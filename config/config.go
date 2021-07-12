package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	AppPort          string
	Profile          string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
}

var Config config

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Config = config{AppPort: os.Getenv("APP_PORT"),
		Profile:          os.Getenv("PROFILE"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabaseHost:     os.Getenv("DATABASE_HOST")}

	if err := GetDBConnection(); err != nil {
		log.Fatal(err)
	}
}
