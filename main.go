package main

import (
	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/routes"
)

func main() {
	routerEngine := gin.Default()

	routes.StartServer(routerEngine)

	if err := routerEngine.Run(); err != nil {
		panic(err)
	}
	// secara default menggunakan port :8080
}
