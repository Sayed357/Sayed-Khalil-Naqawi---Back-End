package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/packages", controllers.CreatePackage)
		api.GET("/packages/:id", controllers.CreatePackage)
		api.POST("/packages", controllers.CreatePackage)
		api.PUT("/packages/:id", controllers.UpdatePackage)
		api.DELETE("/packages/:id", controllers.DeletePackage)
	}
}
