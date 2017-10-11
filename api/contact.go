package api

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/s4kibs4mi/rest-in-go/data"
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

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented",
	})
}

func ListContact(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.Header.Get("user_id"))
	if err == nil {
		contacts := data.GetContactsByUserId(userId)
		json.NewEncoder(w).Encode(Response{
			Code: http.StatusOK,
			Data: contacts,
		})
		return
	}
	json.NewEncoder(w).Encode(Response{
		Code: http.StatusBadRequest,
	})
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotImplemented,
		Message: "Not Implemented",
	})
}
