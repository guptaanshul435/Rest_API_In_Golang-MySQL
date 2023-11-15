package models

type User struct {
	UserId int `json:"userid,omitempty"`
	Name string `json:"name,omitempty"`
	Address string  `json:"address"`
    GmailId  string  `json:"emailid"`
	PhoneNo  string  `json:"number"`
}
