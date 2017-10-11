package data

import (
	"rest-in-go/db"
	"fmt"
)

type User struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

func (u *User) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS users(user_id INTEGER PRIMARY KEY AUTOINCREMENT, user_name TEXT, user_password TEXT);"
	res := db.Exec(query)
	if res {
		query = "INSERT INTO users(user_name, user_password) VALUES('%s', '%s')"
		query = fmt.Sprintf(query, u.UserName, u.UserPassword)
		res = db.Exec(query)
		return res
	}
	return false
}

func (u *User) Delete() {

}

func (u *User) ShouldOpenTheDoor() bool {
	query := "SELECT * FROM users WHERE user_name='%s' AND user_password='%s';"
	query = fmt.Sprintf(query, u.UserName, u.UserPassword)
	rows, err := db.GetRows(query)
	if err != nil {
		return false
	}
	res := rows.Next()
	rows.Close()
	return res
}

func GetUser(userName string) User {
	query := "SELECT * FROM users WHERE user_name='%s';"
	query = fmt.Sprintf(query, userName)
	rows, err := db.GetRows(query)
	if err != nil {
		return User{}
	}
	if !rows.Next() {
		return User{}
	}
	user := User{}
	res := rows.Scan(&user.UserId, &user.UserName, &user.UserPassword)
	defer rows.Close()
	if res != nil {
		return User{}
	}
	return user
}
