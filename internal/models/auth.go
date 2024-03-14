package models

import "time"


type Token struct {
	Token string
	Id string
	Role Role
	Datetime *time.Time
}