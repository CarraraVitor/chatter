package types

import (
	"github.com/gorilla/websocket"
)


type Client struct {
	Chat *Chat
	Conn *websocket.Conn
    User *User
	Send chan SendMessage
}

