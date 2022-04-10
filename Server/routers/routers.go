package routers

import (
	"NeChat/controller"
	"NeChat/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	router.POST("/api/login", controller.LoginHandler)
	router.POST("/api/register", controller.RegisterHandler)
	router.GET("/api/ws", middleware.JwtAuth(), controller.WsHandler)
	router.DELETE("/api/user", middleware.JwtAuth(), controller.DeleteUserHandler)

	return router
}
