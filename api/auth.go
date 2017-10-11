package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/s4kibs4mi/rest-in-go/data"
)

func Authorization(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(r.Header.Get("user_id"))
		if err == nil {
			accessToken := r.Header.Get("access_token")
			if userId > 0 && accessToken != "" {
				session := data.GetSession(accessToken, userId)
				if session.AccessToken == accessToken && session.UserId == userId {
					h.ServeHTTP(w, r)
					return
				}
			}
		}
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request.",
		})
		return
	}
}
