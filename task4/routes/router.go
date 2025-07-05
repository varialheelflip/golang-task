package routes

import (
	"blog_system/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		user := api.Group("/users")
		{
			user.POST("/register", controllers.Register)
			user.POST("/login", controllers.Login)
		}
	}

	{
		post := api.Group("/posts")
		{
			post.POST("", controllers.Create)
			post.GET("/:id", controllers.Detail)
			post.GET("/page", controllers.Page)
			post.PUT("", controllers.Update)
			post.DELETE("/:id", controllers.Delete)
		}
	}
}
