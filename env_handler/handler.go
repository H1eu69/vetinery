package env_handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentHandler struct {
	DbDriver string
	DbSource string
}

func GetEnv() *EnvironmentHandler {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Print("DbDriver:\n")
		log.Print(os.Getenv("DbDriver"))

	return &EnvironmentHandler{
		DbDriver: os.Getenv("DbDriver"),
		DbSource: os.Getenv("DbSource"),
	}
}
