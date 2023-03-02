package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/model"
)

func Login(c *gin.Context) {
	var uc model.UserCredential
	h := model.AuthHeader{}
	h.AuthorizationHeader = "ini_token"
	if err := c.ShouldBindJSON(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":   uc.Username,
			"token": h.AuthorizationHeader,
		})
	}
}
