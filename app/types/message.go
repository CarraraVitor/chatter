package types

import (
	"time"

	"github.com/google/uuid"
)

type RecvMessage struct {
	ChatMessage string `json:"chat_message"`
	Headers     struct {
		HXRequest     string `json:"HX-Request"`
		HXTrigger     string `json:"HX-Trigger"`
		HXTriggerName string `json:"HX-Trigger-Name"`
		HXTarget      string `json:"HX-Target"`
		HXCurrentURL  string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
}

type SendMessage struct {
    Id       uuid.UUID
	Message  string
    Chat     *Chat
	Sender   User
	Receiver User
	SendAt   time.Time
}

type DBMessage struct {
    Id       uuid.UUID
	Message  string
	SendAt   time.Time
    UserId   uuid.UUID
    ChatId   uuid.UUID
}

