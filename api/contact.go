package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"rest-in-go/data"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.Header.Get("user_id"))
	if err == nil {
		contact := data.Contact{}
		err := json.NewDecoder(r.Body).Decode(&contact)
		if err == nil {
			contact.UserId = userId
			if contact.Save() {
				json.NewEncoder(w).Encode(Response{
					Code:    http.StatusOK,
					Message: "Contact successfully created.",
				})
				return
			}
		}
	}
	json.NewEncoder(w).Encode(Response{
		Code: http.StatusBadRequest,
	})
}

func ListContact(w http.ResponseWriter, r *http.Request) {

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {

}
