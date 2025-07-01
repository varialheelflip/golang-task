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
}
