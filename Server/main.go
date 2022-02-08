package main

import (
	"NeChat/config"
	"NeChat/dao"
	"NeChat/logic"
	"NeChat/models"
	"NeChat/routers"
	"fmt"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //release mode on
	dao.InitSQL()
	go logic.WsStart(&models.Manager)

	router := routers.SetupRouter()
	err := router.Run(config.ServerPort)
	if err != nil {
		fmt.Println("[ERR] HTTP server start failed,", err)
		return
	}
}
