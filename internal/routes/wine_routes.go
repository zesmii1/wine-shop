package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wine-shop/internal/auth"
	"wine-shop/internal/delivery"
	"wine-shop/internal/middleware"
	"wine-shop/internal/repository"
	"wine-shop/internal/service"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	wineRepo := repository.NewWineRepository(db)
	userRepo := repository.NewUserRepository(db)

	wineService := service.NewWineService(wineRepo)
	userService := service.NewUserService(userRepo)

	wineHandler := delivery.NewWineHandler(wineService)
	userHandler := delivery.NewUserHandler(userService)

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", auth.Register(db))
		authGroup.POST("/login", auth.Login(db))
	}

	wines := r.Group("/api/v1/wines")
	{
		wines.GET("/", wineHandler.GetAll)
		wines.GET("/:id", wineHandler.GetById)

		wines.POST("/", middleware.AuthRequired(db), middleware.RequireRole("admin"), wineHandler.Create)
		wines.PUT("/:id", middleware.AuthRequired(db), middleware.RequireRole("admin"), wineHandler.Update)
		wines.DELETE("/:id", middleware.AuthRequired(db), middleware.RequireRole("admin"), wineHandler.Delete)
	}

	users := r.Group("/api/users")
	{
		users.GET("/", userHandler.GetAll)
		users.DELETE("/:id", userHandler.Delete)
	}

	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired(db))
	{
		protected.GET("/me", auth.Me(db))
	}

	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middleware.AuthRequired(db), middleware.RequireRole("admin"))
	{
		adminRoutes.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, Admin!"})
		})
	}

}
