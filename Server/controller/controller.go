package controller

import (
	"NeChat/logic"
	"NeChat/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandle(c *gin.Context) {
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

func RegisterHandle(c *gin.Context) {
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

func ChatHandle(c *gin.Context) {

}
