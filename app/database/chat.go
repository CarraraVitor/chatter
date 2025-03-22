package database

import (
	"database/sql"
	"fmt"

	"chatter/app/types"

	"github.com/google/uuid"
)


func ListChats() ([]types.Chat, error) {
    db := Open()
    defer db.Close()

    chats := make([]types.Chat, 0, 10)
    query := "" + 
    "SELECT id, name  " +
    "FROM Chats "
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        chat := types.Chat{}
        rows.Scan(&chat.Id, &chat.Name)
        chats = append(chats, chat)
    }
    return chats, nil
}

func ChatIdsByUser(user_id uuid.UUID) ([]uuid.UUID, error) {
    db := Open()
    defer db.Close()

    query := "" + 
    "SELECT chat_id " + 
    "FROM ChatMembers " + 
    "WHERE user_id = ?"
    rows, err := db.Query(query, user_id)
    if err != nil {
        return nil, err
    }
    ids := []uuid.UUID{}
    for rows.Next(){
        id := uuid.UUID{}
        rows.Scan(&id)
        ids = append(ids, id)
    }
    return ids, nil
}

func GetChatMembers(chat_id uuid.UUID) ([]types.User, error) {
    db := Open()
    defer db.Close()

    query := "" +
    "SELECT id, name, email, profilepic " + 
    "FROM Users " + 
    "WHERE id IN ( " +
    "  SELECT user_id " + 
    "  FROM ChatMembers " + 
    "  WHERE chat_id = ?" + 
    ")"
    rows, err := db.Query(query, chat_id)
    if err != nil {
        return nil, err
    }
    us := []types.User{}
    for rows.Next() {
        u := types.User{}
        rows.Scan(&u.Id, &u.Name, &u.Email, &u.ProfilePic)
        us = append(us, u)
    }
    return us, nil
}

func SaveChat(c types.Chat) (sql.Result, error) {
    db := Open()
    defer db.Close()
    res, err := db.Exec("INSERT INTO Chats (id, name) VALUES (?, ?)", c.Id, c.Name)
    return res, err
}


func DeleteChat(c types.Chat) (sql.Result, error) {
    db := Open()
    defer db.Close()
    res, err := db.Exec("DELETE FROM Chats WHERE id = ?", c.Id)
    return res, err
}

func ChatAddMember(user_id, chat_id uuid.UUID) (sql.Result, error) {
    db := Open()
    defer db.Close()
    id := uuid.New()
    exc := "INSERT INTO ChatMembers (id, user_id, chat_id) VALUES (?, ?, ?);"
    res, err := db.Exec(exc, id, user_id, chat_id)
    return res, err
}

func ChatRemoveMember(user_id, chat_id uuid.UUID) (sql.Result, error) {
    db := Open()
    defer db.Close()
    fmt.Printf("user_id: %s\nchat_id: %s\n", user_id, chat_id)
    del := "DELETE FROM ChatMembers WHERE user_id = ? AND chat_id = ?"
    res, err := db.Exec(del, user_id, chat_id)
    nrows, _ := res.RowsAffected()
    fmt.Printf("remove member:\ndel: %s\nnrows: %d\nerr: %s\n",del, nrows, err)
    return res, err
}


