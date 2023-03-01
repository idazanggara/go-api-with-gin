package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routerEngine := gin.Default()

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})

	if err := routerEngine.Run(); err != nil {
		panic(err)
	}
	// secara default menggunakan port :8080
}
