package router

import (
	handlers "forum/root/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("/", userHandler.Register)
			}
		}
	}

	return r
}
