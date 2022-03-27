package Routes

import (
	"gorm-db/Handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", Handlers.HealthCheck)

	v1 := router.Group("/v1")
	{
		v1.GET("api/verse", Handlers.GetAllVerses)
		v1.POST("api/verse", Handlers.CreateVerse)
		v1.GET("api/category", Handlers.GetAllCategories)

		// v1.GET("todo/:id", Controllers.GetATodo)
		// v1.PUT("todo/:id", Controllers.UpdateATodo)
		// v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
	return router
}
