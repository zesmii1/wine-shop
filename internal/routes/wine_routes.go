package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wine-shop/internal/delivery"
	"wine-shop/internal/repository"
	"wine-shop/internal/service"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	wineRepo := repository.NewWineRepository(db)
	wineService := service.NewWineService(wineRepo)
	wineHandler := delivery.NewWineHandler(wineService)

	wines := r.Group("/api/v1/wines")
	{
		wines.GET("/", wineHandler.GetAll)
		wines.GET("/:id", wineHandler.GetById)
		wines.POST("/", wineHandler.Create)
		wines.PUT("/:id", wineHandler.Update)
		wines.DELETE("/:id", wineHandler.Delete)
	}
}
