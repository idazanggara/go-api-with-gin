package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/controller"
)

func ServerAuth(routerEngine *gin.Engine) *gin.Engine {

	rgAuth := routerEngine.Group("/auth")
	rgAuth.POST("/login", controller.Login)

	return routerEngine

}
