package main

import (
	"fmt"
	"log"

	"gorm-gin-pg/initializers"
	"gorm-gin-pg/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("? Migration failed", err)
	}
	fmt.Println("? Migration complete")
}
