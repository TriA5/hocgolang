package router

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	RegisterRoleRoutes(r)
	RegisterUserRoutes(r)
}
