package database

import (
	"database/sql"

    "chatter/app/types"
)


func GetSessionAndUserBySessionId(sessionid string) (types.Session, types.User, error) {
	db := Open()
	defer db.Close()
    query := "" + 
        "SELECT " + 
        "Sessions.id, Sessions.user_id, Sessions.timestamp, Sessions.expires_at, " + 
        "Users.id, Users.name, Users.email, Users.profilepic " + 
        "FROM Sessions " + 
        "INNER JOIN Users " + 
        "ON Users.id = Sessions.user_id " + 
        "WHERE Sessions.id = ?"
	row := db.QueryRow(query, sessionid)
    var session types.Session
    var user types.User
    err := row.Scan(&session.Id, &session.UserId, &session.Timestamp, &session.ExpiresAt, &user.Id, &user.Name, &user.Email, &user.ProfilePic)
    if err != nil {
        return types.Session{}, types.User{}, err
    }
	return session, user, nil
}

func AddSession(session types.Session) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("INSERT INTO Sessions (id, user_id, timestamp, expires_at) VALUES (?, ?, ?, ?)", session.Id, session.UserId, session.Timestamp, session.ExpiresAt)
	return res, err
}

func UpdateSession(session types.Session) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("UPDATE Sessions SET user_id = ?, expires_at = ? WHERE id = ?", session.UserId, session.ExpiresAt, session.Id)
	return res, err
}

func DeleteSession(sessionid string) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("DELETE FROM Sessions WHERE id = ?", sessionid)
	return res, err

}
