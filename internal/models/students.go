package models


type Students struct{
	ID string  `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Score int `json:"score"`
}