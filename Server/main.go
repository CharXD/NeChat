package main

import (
	"NeChat/config"
	"NeChat/dao"
	"NeChat/routers"
	"fmt"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //release mode on
	dao.InitSQL()

	router := routers.SetupRouter()
	err := router.Run(config.ServerPort)
	if err != nil {
		fmt.Println("[ERR] HTTP server start failed,", err)
		return
	}
}
