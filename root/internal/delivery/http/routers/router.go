package router

import (
	handlers "forum/root/internal/delivery/http/handlers"
	middleware "forum/root/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, postHandler *handlers.PostHandler, jwtSecretKey string) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			auth := v1.Group("/auth")
			{
				auth.POST("/register", authHandler.Register)
				auth.POST("/login", authHandler.Login)
			}

			user := v1.Group("/user")
			{
				user.GET("/", middleware.JWTMiddleware(jwtSecretKey), userHandler.GetUser)
			}

			post := v1.Group("/post")
			{
				post.POST("/", middleware.JWTMiddleware(jwtSecretKey), postHandler.CreatePost)
				post.GET("/:id", middleware.JWTMiddleware(jwtSecretKey), postHandler.GetPost)
				post.PUT("/:id", middleware.JWTMiddleware(jwtSecretKey), postHandler.UpdatePost)
				post.DELETE("/:id", middleware.JWTMiddleware(jwtSecretKey), postHandler.DeletePost)
			}
		}
	}

	return r
}
