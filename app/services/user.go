package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"chatter/app/database"
	"chatter/app/types"
)


func NewUser(name, email, password string) (types.User, error) {
    hashed, err := HashPassword(password)
    if err != nil {
        log.Printf("[ERROR] NewUser: HashPassword: %s\n", err.Error())
        err = fmt.Errorf("NewUser: failed to hash password: %s", err)
        return types.User{}, err
    }

    user := types.User{
        Id:       uuid.New(),
        Name:     name,
        Email:    email,
        Password: hashed,
        ProfilePic: "default-pfp.png",
    }
    return user, nil
}

func SaveUser(user types.User) (types.User, error) {
    _, err := database.AddUser(user)
    return user, err
}

func AnonUser() types.User {
	uid := uuid.New()
	name := fmt.Sprintf("user_%s", uid.String())
    user := types.User{
		Id:       uid,
		Name:     name,
		Email:    fmt.Sprintf("%s@anon", uid.String()),
		Password: uid.String(),
        ProfilePic: "default-pfp.png",
	}
    return user
}

type IsLogged = bool
func UserFromSessionCookie(r *http.Request) (types.User, IsLogged) {
	ck, err := r.Cookie("SessionToken")
	if err != nil {
        log.Printf("user from cookie: get session token: %s\n", err)
		return AnonUser(), false
	}
	res, err := ValidateSessionToken(ck.Value)
	if err != nil {
        log.Printf("user from cookie: validate session token: %s\n", err)
		return AnonUser(), false
	}
	return res.User, true
}

func UserContext(r *http.Request) context.Context {
	user, _ := UserFromSessionCookie(r)
	ctx := context.WithValue(context.Background(), "user", user)
	return ctx
}

func IsAnonUser(user types.User) bool {
    email := user.Email
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return true
	}
	tail := parts[1]
	tail_parts := strings.Split(tail, ".")
	if len(tail_parts) < 1 {
		return true
	}
	domain := tail_parts[0]
    if domain == "anon" {
        return true
    }
    return false 
}

func UserSearchEmail(search string) []types.User {
    usrs, err := database.GetUserMatchEmail(search)
    if err != nil {
        log.Printf("user search: %s\n", err)
        return nil
    }
    return usrs

}
