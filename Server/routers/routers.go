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

	router.POST("/login", controller.LoginHandle)
	router.POST("/register", controller.RegisterHandle)
	router.POST("/chat", middleware.JwtAuth(), controller.ChatHandle)

	return router
}
