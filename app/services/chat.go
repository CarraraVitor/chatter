package services

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"chatter/app/database"
	"chatter/app/types"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Chats map[uuid.UUID]*types.Chat

func init() {
	Chats = make(map[uuid.UUID]*types.Chat)
	cs, err := database.ListChats()
	if err != nil {
		panic(fmt.Sprintf("init chat.go: load Chats: %s", err))
	}
    for _, c := range cs {
        c.Init()
        Chats[c.Id] = &c
        go c.Run()
    }
}


func NewChat(name string) *types.Chat {
	return &types.Chat{
		Id:         uuid.New(),
		Name:       name,
		Clients:    make(map[*types.Client]bool),
		Broadcast:  make(chan types.SendMessage),
		Register:   make(chan *types.Client),
		Unregister: make(chan *types.Client),
	}
}

func ListChatsWithUser(r *http.Request) ([]types.Chat, error) {
    user, _ := UserFromSessionCookie(r)
    user_rooms, err := ChatIdsByUser(user.Id)
    if err != nil {
        return nil, err
    }
	rooms := make([]types.Chat, 0, len(Chats))
	for _, chat := range Chats {
        if _, ok := user_rooms[chat.Id]; ok {
            rooms = append(rooms, *chat)
        }
	}
    return rooms, nil
}

func ChatIdsByUser(user_id uuid.UUID) (map[uuid.UUID]bool, error) {
    ids, err := database.ChatIdsByUser(user_id) 
    if err != nil {
        return nil, err
    }
    res := make(map[uuid.UUID]bool)
    for _, id := range ids {
        res[id] = true
    }
    return res, nil
}

func ChatMembers(chat_id uuid.UUID) []types.User {
    ms, err := database.GetChatMembers(chat_id)
    if err != nil {
        log.Printf("chat members: %s\n", err)
        return []types.User{}
    }
    return ms
}

func ChatAddMember(usrid, chtid string) error { 
    user_id, err := uuid.Parse(usrid) 
    if err != nil { return err }
    chat_id, err := uuid.Parse(chtid) 
    if err != nil { return err }
    _, err = database.ChatAddMember(user_id, chat_id)
    return err
}

func ChatRemoveMember(usrid, chtid string) error {
    user_id, err := uuid.Parse(usrid) 
    if err != nil { return err }
    chat_id, err := uuid.Parse(chtid) 
    if err != nil { return err }
    _, err = database.ChatRemoveMember(user_id, chat_id)
    return err
}

func DeleteChat(c *types.Chat) {
    close(c.Broadcast)
    close(c.Register)
    for client := range c.Clients {
        c.Unregister <- client
    }
    close(c.Unregister)
    database.DeleteChat(*c)
	delete(Chats, c.Id)
}


func ConnectNewClient(w http.ResponseWriter, r *http.Request ,id uuid.UUID, user types.User) error {
	c, ok := Chats[id]
	if !ok {
        return fmt.Errorf("connect new client: failed retrieving chatroom with id %s", id)
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
        return fmt.Errorf("connect new client: failed to upgrade connection:  %s", err)
	}
	if conn == nil {
        return fmt.Errorf("connect new client: nil connection")
	}

	client := &types.Client{Chat: c, Conn: conn, Send: make(chan types.SendMessage, 256), User: &user}
	c.Register <- client
	go readClient(client)
	go writeClient(client) 
    return nil
}
