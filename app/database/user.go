package database

import (
	"database/sql"
	"errors"
	"log"

	"chatter/app/types"
)

var UserNotFound error = errors.New("User Not Found")

func GetUserByEmail(email string) (types.User, error) {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT id, name, email, password, profilepic FROM Users WHERE email = ?", email)
	if err != nil {
		log.Printf("[ERROR] :%s", err.Error())
		return types.User{}, err
	}
	var users []types.User
	for rows.Next() {
		var user types.User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.ProfilePic)
		users = append(users, user)
	}
	if len(users) < 1 {
		return types.User{}, UserNotFound
	}
	return users[0], nil
}

func AddUser(user types.User) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("INSERT INTO Users (id, name, email, password, profilepic) VALUES (?, ?, ?, ?, ?)", user.Id.String(), user.Name, user.Email, user.Password, user.ProfilePic)
	return res, err
}

func UpdateUser(user types.User) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("UPDATE Users SET name = ?, email = ?, password = ?, profilepic = ? WHERE id = ?", user.Name, user.Email, user.Password, user.ProfilePic, user.Id)
	return res, err
}

func GetUserMatchEmail(search string) ([]types.User, error) {
	db := Open()
	defer db.Close()

	query := "" +
		"SELECT id, name, email, profilepic " +
		"FROM Users " +
		"WHERE email LIKE '%' || ? || '%'"
	rows, err := db.Query(query, search)
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
