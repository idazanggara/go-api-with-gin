package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	routerEngine := gin.Default()

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})

	routerEngine.GET("/greeting/:name", greeting)

	routerEngine.POST("/login", login)

	if err := routerEngine.Run(); err != nil {
		panic(err)
	}
	// secara default menggunakan port :8080
}

func greeting(c *gin.Context) {
	name := c.Param("name")
	kec := c.Query("kecamatan")
	kel := c.Query("kelurahan")
	c.String(http.StatusOK, "Hello %s saat ini kamu berada di kec %s kel %s", name, kec, kel)
}

func login(c *gin.Context) {
	// username := c.PostForm("username")
	// c.PostForm("password")
	// c.String(http.StatusOK, "Hello %s", username)
	// func login(response, request)
	var uc UserCredential
	if err := c.ShouldBind(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": uc.Username,
		})
	}

}
