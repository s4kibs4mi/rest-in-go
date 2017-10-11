package data

type Contact struct {
	ContactId     int    `json:"contact_id"`
	UserId        int    `json:"user_id"`
	ContactName   string `json:"contact_name"`
	ContactNumber string `json:"contact_number"`
}

func (c *Contact) Save() {

}

func (c *Contact) Delete() {

}

func GetByContactId(contactId int) Contact {
	return Contact{}
}

func GetByUserId(userId int) []User {
	return []User{}
}
