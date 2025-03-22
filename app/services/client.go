package services

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
    "chatter/app/types"
)

const (
	WriteWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)


func readClient(c *types.Client) {
	defer func() {
		c.Chat.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client.Read: %v", err)
			}
			break
		}
        tosend := procSaveMsg(msg, *c)
		c.Chat.Broadcast <- tosend
	}
}

func writeClient(c *types.Client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
            msg.Receiver = *c.User
			c.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				log.Printf("error in Client.Write, closing message.\n")
                log.Printf("client: %+v\n", c)
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("error in Client.Write: %s\n", err)
				return
			}

			renderSendMessage(w, msg)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				msg, ok := <-c.Send
				if ok {
					renderSendMessage(w, msg)
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
