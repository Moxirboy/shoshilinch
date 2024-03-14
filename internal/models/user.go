package models

type Role string

const (
	Student Role=Role("student")
	Teacher Role=Role("teacher")
	Admin  	Role=Role("admin")
)

type User struct {
	ID          string 
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	Role   		Role
}

