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
			var userController = controllers.UserController{}
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
		}
	}

	{
		post := api.Group("/posts")
		{
			var postController = controllers.PostController{}
			post.POST("", postController.Create)
			post.GET("/:id", postController.Detail)
			post.GET("/page", postController.Page)
			post.PUT("", postController.Update)
			post.DELETE("/:id", postController.Delete)
		}
	}

	{
		post := api.Group("/comments")
		{
			var commentController = controllers.CommentController{}
			post.POST("", commentController.Create)
			post.GET("/:postId", commentController.List)
		}
	}
}
