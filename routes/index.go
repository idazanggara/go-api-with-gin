package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/controller"
	"github.com/idazanggara/go-api-with-gin/middleware"
)

func StartServer(app *gin.Engine) *gin.Engine {

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})

	app.Use(middleware.AuthMiddleware()) // implement middleware di global

	// group routing
	// group login
	ServerAuth(app)
	// rgAuth := app.Group("/auth")
	// rgAuth.POST("/login", controller.Login)

	// group cars
	rgMaster := app.Group("/master")
	rgMaster.POST("/cars", controller.CreateCar)
	// /cars => rgMaster.GET(data)
	rgMaster.GET("/cars", controller.GetAllCar)
	rgMaster.PUT("/cars/:carID", controller.UpdateCar)
	rgMaster.GET("/cars/:carID", controller.GetCarById)
	rgMaster.DELETE("/cars/:carID", controller.DeleteCar)

	return app
}
