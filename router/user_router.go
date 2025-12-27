package router

import (
	"be-ep/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("", handler.CreateUser)
		user.GET("", handler.GetUsers)
		user.GET("/:id", handler.GetUserByID)
		user.PUT("/:id", handler.UpdateUser)
		user.DELETE("/:id", handler.DeleteUser)
	}
}
