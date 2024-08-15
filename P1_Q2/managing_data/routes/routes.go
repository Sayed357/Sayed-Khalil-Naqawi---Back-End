package routes

import (
	"managing_data/controllers"
	"managing_data/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{

		secured := v1.Group("/")
		secured.Use(middlewares.AuthMiddleware())
		secured.Use(middlewares.RateLimiterMiddleware())
		{
			secured.GET("/posts", controllers.GetPosts)
			secured.GET("/posts/:id", controllers.GetPost)
			secured.POST("/posts", controllers.CreatePost)
			secured.PUT("/posts/:id", controllers.UpdatePost)
			secured.DELETE("/posts/:id", controllers.DeletePost)
		}
	}
}
