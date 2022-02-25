package controller

import (
	"NeChat/logic"
	"NeChat/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	userInfo := models.User{}
	resp, err := c.GetRawData()
	//创建用户对象、接收服务端json
	if err != nil {
		fmt.Println("[ERROR]Response get failed,", err)
	}
	err = json.Unmarshal(resp, &userInfo)
	if err != nil {
		fmt.Println("[ERROR]Json unmarshal failed,", err)
	}

	c.JSON(http.StatusOK, logic.Login(&userInfo))
}

func RegisterHandler(c *gin.Context) {
	userInfo := models.User{}
	resp, err := c.GetRawData()
	if err != nil {
		fmt.Println("[ERROR]Response get failed,", err)
	}
	err = json.Unmarshal(resp, &userInfo)
	if err != nil {
		fmt.Println("[ERROR]Json unmarshal failed,", err)
	}

	c.JSON(http.StatusOK, logic.Register(&userInfo))
}

// WsHandler 升级协议,自定义信息等
func WsHandler(c *gin.Context) {
	username := c.Query("username")
	friendName := c.Query("friend_id")
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &models.Client{
		ID:     logic.CreatId(username, friendName),
		Socket: conn,
		Send:   make(chan []byte),
	}
	models.Manager.Register <- client
	go logic.Read(client)
	go logic.Write(client)
}

func DeleteUserHandler(c *gin.Context) {
	username, err := c.Get("Username")
	if err {
		fmt.Println("[ERROR]Handler var get failed.")
	}
	c.JSON(http.StatusOK, logic.DeleteUser(username.(string)))
}
