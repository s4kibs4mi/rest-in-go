package api

import "time"

type Session struct {
	SessionId    int       `json:"session_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpireTime   time.Time `json:"expire_time"`
}

func (s *Session) Save() {

}

func (s *Session) Delete() {

}

func GetSession() Session {
	return Session{}
}
