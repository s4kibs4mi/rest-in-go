package data

import (
	"fmt"
	"github.com/s4kibs4mi/rest-in-go/db"
)

type Contact struct {
	ContactId     int    `json:"contact_id"`
	UserId        int    `json:"-"`
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

func GetContactsByUserId(userId int) []Contact {
	query := "SELECT * FROM contacts WHERE user_id=%d"
	query = fmt.Sprintf(query, userId)
	rows, err := db.GetRows(query)
	if err != nil {
		return []Contact{}
	}
	defer rows.Close()
	contacts := []Contact{}
	for rows.Next() {
		contact := Contact{}
		rows.Scan(&contact.ContactId, &contact.UserId, &contact.ContactName, &contact.ContactNumber)
		contacts = append(contacts, contact)
	}
	return contacts
}
