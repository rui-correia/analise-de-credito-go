package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
}

func Ping(ctx *gin.Context) {

	ctx.JSONP(http.StatusOK, "pong")
}
