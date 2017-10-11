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
	query := "CREATE TABLE IF NOT EXISTS users(userId INTEGER PRIMARY KEY AUTOINCREMENT, userName TEXT, userPassword TEXT);"
	res := db.Exec(query)
	if res {
		query = "INSERT INTO users(userName, userPassword) VALUES('%s', '%s')"
		query = fmt.Sprintf(query, u.UserName, u.UserPassword)
		res = db.Exec(query)
		return res
	}
	return false
}

func (u *User) Delete() {

}

func (u *User) ShouldOpenTheDoor() bool {
	query := "SELECT * FROM users WHERE userName='%s' AND userPassword='%s';"
	query = fmt.Sprintf(query, u.UserName, u.UserPassword)
	rows, err := db.GetRows(query)
	if err != nil {
		return false
	}
	res := rows.Next()
	rows.Close()
	return res
}

func GetUser(userId int) User {
	return User{}
}
