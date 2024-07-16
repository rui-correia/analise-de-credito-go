package server

import (
	"analise-de-credito-go/controller"
	"github.com/gin-gonic/gin"
)

func mapRoutes() *gin.Engine {
	router := gin.Default()
	mapControllerRoutes(router)
	return router
}

func mapControllerRoutes(router *gin.Engine) {
	controller.MapRoutes(router)
}
