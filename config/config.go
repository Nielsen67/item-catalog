package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// Get PORT with default "8070"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
	}

	// Get ENV with default "dev"
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	return &Config{
		Port: port,
		Env:  env,
	}
}
