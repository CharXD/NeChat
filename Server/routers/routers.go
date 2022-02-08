package routers

import (
	"NeChat/controller"
	"NeChat/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/login", controller.LoginHandler)
	router.POST("/register", controller.RegisterHandler)
	router.GET("/ws", middleware.JwtAuth(), controller.WsHandler)

	return router
}
