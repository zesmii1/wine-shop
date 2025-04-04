package routes

import (
	"github.com/gin-gonic/gin"
	"wine-shop/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	wineGroup := r.Group("/wines")
	{
		wineGroup.GET("", handlers.GetWines)
		wineGroup.GET("/:id", handlers.GetWineByID)
		wineGroup.POST("", handlers.CreateWine)
		wineGroup.PUT("/:id", handlers.UpdateWine)
		wineGroup.DELETE("/:id", handlers.DeleteWine)
	}

	return r
}
