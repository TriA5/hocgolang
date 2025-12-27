package router

import (
	"be-ep/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.Engine) {
	role := r.Group("/roles")
	{
		role.POST("", handler.CreateRole)
		role.GET("", handler.GetRoles)
		role.GET("/:id", handler.GetRoleByID)
		role.PUT("/:id", handler.UpdateRole)
		role.DELETE("/:id", handler.DeleteRole)
	}
}
