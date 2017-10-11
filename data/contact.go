package data

import (
	"rest-in-go/db"
	"fmt"
)

type Contact struct {
	ContactId     int    `json:"contact_id"`
	UserId        int    `json:"user_id"`
	ContactName   string `json:"contact_name"`
	ContactNumber string `json:"contact_number"`
}

func (c *Contact) Save() bool {
	query := "CREATE TABLE IF NOT EXISTS contacts(contact_id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER, contact_name TEXT, contact_number TEXT);"
	if db.Exec(query) {
		query = "INSERT INTO contacts(user_id, contact_name, contact_number) VALUES(%d, '%s', '%s');"
		query = fmt.Sprintf(query, c.UserId, c.ContactName, c.ContactNumber)
		return db.Exec(query)
	}
	return false
}

func (c *Contact) Delete() {

}

func GetByContactId(contactId int) Contact {
	return Contact{}
}

func GetByUserId(userId int) []User {
	return []User{}
}
