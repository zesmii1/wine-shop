package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wine-shop/internal/models"
	"wine-shop/internal/routes"
)

func main() {
	dsn := "postgres://myuser:mypassword@localhost:5433/mydatabase?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = db.AutoMigrate(&models.Wine{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db) // Передаем db в SetupRoutes
	r.Run(":8080")
}
