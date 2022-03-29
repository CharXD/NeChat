package logic

import (
	"NeChat/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// WsStart is  项目运行前, 协程开启start -> go Manager.WsStart()
func WsStart(manager *models.ClientManager) {
	for {
		log.Println("<---管道通信--->")
		select {
		case conn := <-models.Manager.Register:
			log.Printf("新用户加入:%v", conn.ID)
			models.Manager.Clients[conn.ID] = conn
			jsonMessage, _ := json.Marshal(&models.Message{Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage
		case conn := <-models.Manager.Unregister:
			log.Printf("用户离开:%v", conn.ID)
			if _, ok := models.Manager.Clients[conn.ID]; ok {
				jsonMessage, _ := json.Marshal(&models.Message{Content: "A socket has disconnected"})
				conn.Send <- jsonMessage
				close(conn.Send)
				delete(models.Manager.Clients, conn.ID)
			}
		case message := <-models.Manager.Broadcast:
			MessageStruct := models.Message{}
			err := json.Unmarshal(message, &MessageStruct)
			if err != nil {
				fmt.Println("[ERROR]Data unmarshal failed.")
			}
			models.AddMessage(&MessageStruct)
			for id, conn := range models.Manager.Clients {
				if id != CreatId(MessageStruct.Recipient, MessageStruct.Sender) {
					continue
				}
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(models.Manager.Clients, conn.ID)
				}
			}
		}
	}
}
func CreatId(uid, fid string) string {
	return uid + "_" + fid
}
func Read(c *models.Client) {
	defer func() {
		models.Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			models.Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		log.Printf("读取到客户端的信息:%s", string(message))
		models.Manager.Broadcast <- message
	}
}

func Write(c *models.Client) {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("发送到到客户端的信息:%s", string(message))

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
