package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idazanggara/go-api-with-gin/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// di cek. endpoint yg di hit apalah "/auth/login"? kalau bener dia disuruh login dulu
		if ctx.Request.URL.Path == "/auth/login" {
			ctx.Next()
		} else {
			h := model.AuthHeader{}

			if err := ctx.ShouldBindHeader(&h); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				ctx.Abort()
			}

			// Selain endpoint ini "/auth/login", di cek, apakah ada token atau enggak?
			if h.AuthorizationHeader == "ini_token" {
				ctx.Next()
				// kalau bener ada di next, bisa nampilin data
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				// kalau enggak ada token, di kasih pesan, invalid token
				ctx.Abort()
			}
		}
	}
}
