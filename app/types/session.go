package types

import (
	"github.com/google/uuid"
)

type Session struct {
	Id        string
	UserId    uuid.UUID
	Timestamp int64
	ExpiresAt int64
}

type SessionValidationResult struct {
	Session Session
	User    User
}
