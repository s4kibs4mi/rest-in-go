package data

import (
	"rest-in-go/db"
	"fmt"
)

type Session struct {
	SessionId    int    `json:"session_id,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserId       int    `json:"user_id,omitempty"`
}

func (s *Session) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS sessions(sessionId INTEGER PRIMARY KEY AUTOINCREMENT, accessToken TEXT, refreshToken TEXT);"
	if db.Exec(query) {
		query = "INSERT INTO sessions(accessToken, refreshToken) VALUES('%s', '%s');"
		query = fmt.Sprintf(query, s.AccessToken, s.RefreshToken)
		return db.Exec(query)
	}
	return false
}

func (s *Session) Delete() {

}

func GetSession() Session {
	return Session{}
}
