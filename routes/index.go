package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/controller"
)

func StartServer() *gin.Engine {
	routerEngine := gin.Default()

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})
	// group routing
	// group login
	rgAuth := routerEngine.Group("/auth")
	rgAuth.POST("/login", controller.Login)

	// group cars
	rgMaster := routerEngine.Group("/master")
	rgMaster.POST("/cars", controller.CreateCar)
	rgMaster.GET("/cars", controller.GetAllCar)
	rgMaster.PUT("/cars/:carID", controller.UpdateCar)
	rgMaster.GET("/cars/:carID", controller.GetCarById)
	rgMaster.DELETE("/cars/:carID", controller.DeleteCar)

	return routerEngine
}
