package services

import (
	"context"
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"time"

	"chatter/app/components"
	"chatter/app/database"
	"chatter/app/types"

	"github.com/google/uuid"
)

func renderSendMessage(w io.Writer, msg types.SendMessage) {
	ctx := context.Background()
	err := components.SendMessage(msg).Render(ctx, w)
	if err != nil {
		log.Printf("msg.Render: %s", err)
	}
}

func procSaveMsg(msg []byte, cln types.Client) types.SendMessage {
	proc := procMsg(msg, cln)
	saveMessage(proc)
	return proc
}

func procMsg(msg []byte, cln types.Client) types.SendMessage {
	pmsg := &types.RecvMessage{}
	err := json.Unmarshal(msg, pmsg)
	if err != nil {
		log.Printf("proccess message: %s\n", err)
		return types.SendMessage{}
	}
	now := time.Now().UTC()
	return types.SendMessage{
		Id:      uuid.New(),
		Message: pmsg.ChatMessage,
		Chat:    cln.Chat,
		Sender:  *(cln.User),
		SendAt:  now,
	}
}

func saveMessage(msg types.SendMessage) {
	_, err := database.SaveMessage(msg)
	if err != nil {
		log.Printf("save message: %s", err)
	}
}

func ListChatMessages(chat *types.Chat) []types.SendMessage {
	return database.ListChatMessages(chat)
}
