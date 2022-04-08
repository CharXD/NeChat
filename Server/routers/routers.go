package routers

import (
	"NeChat/controller"
	"NeChat/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	router.POST("/login", controller.LoginHandler)
	router.POST("/register", controller.RegisterHandler)
	router.GET("/ws", middleware.JwtAuth(), controller.WsHandler)
	router.DELETE("/user", middleware.JwtAuth(), controller.DeleteUserHandler)

	return router
}
