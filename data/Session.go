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
	query := "CREATE TABLE IF NOT EXISTS sessions(session_id INTEGER PRIMARY KEY AUTOINCREMENT, access_token TEXT, refresh_token TEXT, user_id INTEGER);"
	if db.Exec(query) {
		query = "INSERT INTO sessions(access_token, refresh_token, user_id) VALUES('%s', '%s', %d);"
		query = fmt.Sprintf(query, s.AccessToken, s.RefreshToken, s.UserId)
		return db.Exec(query)
	}
	return false
}

func GetSession(accessToken string, userId int) Session {
	query := "SELECT * FROM sessions WHERE access_token='%s' AND user_id=%d"
	query = fmt.Sprintf(query, accessToken, userId)
	rows, err := db.GetRows(query)
	session := Session{}
	if err == nil {
		if rows.Next() {
			rows.Scan(&session.SessionId, &session.AccessToken, &session.RefreshToken, &session.UserId)
			rows.Close()
		}
	}
	return session
}
