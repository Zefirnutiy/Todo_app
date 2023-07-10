package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//эту реализацию нужно пересмотреть
func CheckToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header["Authorization"]
	
	if authHeader == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message":"не авторизован",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	auth := strings.Split(authHeader[0], " ")

	if len(auth) != 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message":"не авторизован",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if auth[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message":"не авторизован",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if _, err := ParseToken(auth[1]); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message":"не авторизован",
		})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}