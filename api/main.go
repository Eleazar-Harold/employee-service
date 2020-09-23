package main

import (
	"log"

	"github.com/eleazar-harold/employee-service/api/app"
	"github.com/eleazar-harold/employee-service/api/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
}
