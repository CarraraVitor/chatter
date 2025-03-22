package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"chatter/app/database"
	"chatter/app/types"
)


func GenerateSessionToken() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	sessionid := base32.StdEncoding.EncodeToString(bytes)
	return sessionid
}

func CreateSession(token string, userid uuid.UUID) types.Session {
	hash := sha256.New()
	hash.Write([]byte(token))
	src := hash.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	sessionid := string(dst)
	session := types.Session{
		Id:        string(sessionid),
		UserId:    userid,
		Timestamp: time.Now().UTC().Unix(),
		ExpiresAt: time.Now().UTC().Unix() + int64(30*24*60*60),
	}
    _, err := database.AddSession(session)
    if err != nil {
        log.Printf("create session: add session to db: %s\n", err)
        return types.Session{}
    }
	return session
}

func ValidateSessionToken(token string) (types.SessionValidationResult, error) {
	hash := sha256.New()
	hash.Write([]byte(token))
	src := hash.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	sessionid := string(dst)
	session, user, err := database.GetSessionAndUserBySessionId(sessionid)
	if err != nil {
		return types.SessionValidationResult{
			Session: types.Session{},
			User:    types.User{},
		}, err
	}

	if time.Now().UTC().Unix() >= session.ExpiresAt {
		database.DeleteSession(sessionid)
		return types.SessionValidationResult{
			Session: types.Session{},
			User:    types.User{},
		}, err
	}

	if time.Now().UTC().Unix() >= (session.ExpiresAt - 15*24*60*60) {
		session.ExpiresAt = time.Now().UTC().Unix() + 30*24*60*60
		database.UpdateSession(session)
	}

	return types.SessionValidationResult{
		Session: session,
		User:    user,
	}, nil
}

func InvalidateSession(sessionid string) {
	database.DeleteSession(sessionid)
}

func SetSessionTokenCookie(w http.ResponseWriter, token string, session types.Session) {
	session_cookie := http.Cookie{
		Name:     "SessionToken",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(session.ExpiresAt, 10),
		Path:     "/",
	}
	http.SetCookie(w, &session_cookie)
}

func DeleteSessionTokenCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "SessionToken",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   0,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
