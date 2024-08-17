package router

import (
	handlers "forum/root/internal/delivery/http/handlers"
	middleware "forum/root/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, postHandler *handlers.PostHandler, jwtSecretKey string) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("/", userHandler.Register)
				user.POST("/login", userHandler.Login)
				user.GET("/", middleware.JWTMiddleware(jwtSecretKey), userHandler.GetUser)
			}

			post := v1.Group("/post")
			{
				post.POST("/", postHandler.CreatePost)
				post.GET("/:id", postHandler.GetPost)
				post.PUT("/:id", postHandler.UpdatePost)
				post.DELETE("/:id", postHandler.DeletePost)
			}
		}
	}

	return r
}
