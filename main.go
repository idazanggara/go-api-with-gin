package main

import "github.com/idazanggara/go-api-with-gin/routes"

func main() {

	// routes.StartServer() == routerEngine

	if err := routes.StartServer().Run(); err != nil {
		panic(err)
	}
	// secara default menggunakan port :8080
}
