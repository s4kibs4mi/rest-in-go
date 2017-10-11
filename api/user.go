package api

import (
	"net/http"
	"encoding/json"
	"github.com/s4kibs4mi/rest-in-go/data"
	"github.com/s4kibs4mi/rest-in-go/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		if user.UserName != "" && user.UserPassword != "" {
			if user.Save() {
				json.NewEncoder(w).Encode(Response{
					Code:    http.StatusOK,
					Message: "User successfully created.",
				})
				return
			} else {
				json.NewEncoder(w).Encode(Response{
					Code:    http.StatusInternalServerError,
					Message: "Something went wrong.",
				})
				return
			}
		}
	}
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusBadRequest,
		Message: "Malformed data received.",
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := data.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		if user.UserName != "" && user.UserPassword != "" {
			if user.ShouldOpenTheDoor() {
				accessToken := utils.GetUUI()
				refreshToken := utils.GetUUI()
				if accessToken != "" && refreshToken != "" {
					session := data.Session{
						AccessToken:  accessToken,
						RefreshToken: refreshToken,
					}
					user = data.GetUser(user.UserName)
					session.UserId = user.UserId
					if session.Save() {
						json.NewEncoder(w).Encode(Response{
							Code: http.StatusOK,
							Data: session,
						})
						return
					}
				}
			} else {
				json.NewEncoder(w).Encode(Response{
					Code:    http.StatusNonAuthoritativeInfo,
					Message: "Username & Password mismatch.",
				})
				return
			}
		}
	}
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusBadRequest,
		Message: "Malformed data received.",
	})
}

func DeleteAccessToken(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented",
	})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented",
	})
}
