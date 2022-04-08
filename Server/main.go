package main

import (
	"NeChat/config"
	"NeChat/dao"
	"NeChat/logic"
	"NeChat/routers"
	"fmt"
)

func main() {
	config.InitConfig()
	//gin.SetMode(gin.ReleaseMode) //release mode on
	dao.InitSQL()
	go logic.WsStart()

	router := routers.SetupRouter()
	err := router.Run(":" + config.ServerConfig.Server.Port)
	if err != nil {
		fmt.Println("[ERR] HTTP server start failed,", err)
		return
	}
}
