package models

import (
	"NeChat/dao"
	"fmt"
	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is return msg
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Date      int    `json:"date,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

func AddMessage(message *Message) {
	sqlStr := "INSERT INTO messages (sender, recipient, content, date) VALUES (?, ?, ?, ?)"
	_, err := dao.DB.Exec(sqlStr, message.Sender, message.Recipient, message.Content, message.Date)
	if err != nil {
		fmt.Println("[ERROR]Inset data(message) failed,", err)
	}
}
