package main

import (
	"github.com/gin-gonic/gin"
	"wine-shop/internal/db"
	"wine-shop/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db.DB)
	r.Static("/frontend", "./frontend")
	r.Run(":8081")

}
