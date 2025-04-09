package types

import (
	"github.com/google/uuid"
)

type Chat struct {
	Id         uuid.UUID
	Name       string
	Clients    map[*Client]bool
	Destroy    chan bool
	Broadcast  chan SendMessage
	Register   chan *Client
	Unregister chan *Client
    Members    []User
}

func (c *Chat) Init() {
	c.Clients = make(map[*Client]bool)
	c.Broadcast = make(chan SendMessage)
	c.Destroy = make(chan bool)
	c.Register = make(chan *Client)
	c.Unregister = make(chan *Client)
}


func (c *Chat) Run() {
	for {
		select {
		case msg := <-c.Broadcast:
			for client := range c.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(c.Clients, client)
				}
			}
		case client := <-c.Register:
			c.Clients[client] = true
		case client := <-c.Unregister:
			if _, ok := c.Clients[client]; ok {
				delete(c.Clients, client)
				close(client.Send)
			}
		case <-c.Destroy:
			goto exit;
		}
	}
exit:
	close(c.Destroy)
	return
}
