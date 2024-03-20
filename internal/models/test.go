package models

type Tests struct{
	Test Test
	Choices []*Choice
}





type Test struct {
	Id string
	UserID string
	Question string
}

