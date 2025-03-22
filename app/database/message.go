package database

import (
	"chatter/app/types"
	"database/sql"
	"log"
)

func SaveMessage(msg types.SendMessage) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec(
		"INSERT INTO Messages (id, content, sent_at, user_id, chat_id) VALUES (?, ?, ?, ?, ?)",
		msg.Id.String(), msg.Message, msg.SendAt, msg.Sender.Id, msg.Chat.Id,
	)
	return res, err
}

func ListChatMessages(chat *types.Chat) []types.SendMessage {
	db := Open()
	defer db.Close()

	query := "" +
		"SELECT " +
		"Messages.id, Messages.content, Messages.sent_at, " +
		"Users.id, Users.name, Users.email, Users.profilepic " +
		"FROM Messages " +
		"LEFT JOIN Users " +
		"ON Messages.user_id = Users.id " +
		"WHERE Messages.chat_id = ?"
	rows, err := db.Query(query, (*chat).Id.String())
	if err != nil {
        log.Printf("list chat messages: query db: %s\n", err)
		return []types.SendMessage{}
	}

	msgs := []types.SendMessage{}
	for rows.Next() {
		msg := types.SendMessage{}
        err := rows.Scan(&msg.Id, &msg.Message, &msg.SendAt, &msg.Sender.Id, &msg.Sender.Name, &msg.Sender.Email, &msg.Sender.ProfilePic)
        if err != nil {
            log.Printf("list chat messages: scan message and user: %s\n", err)
            continue
        }
        msg.Chat = chat
		msgs = append(msgs, msg)
	}

	return msgs
}
