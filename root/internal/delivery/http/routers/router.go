package router

import (
	handlers "forum/root/internal/delivery/http/handlers"
	middleware "forum/root/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, jwtSecretKey string) *gin.Engine {
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
		}
	}

	return r
}
