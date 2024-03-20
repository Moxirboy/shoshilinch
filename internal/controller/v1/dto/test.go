package dto


type Tests struct{
	Test Test
	Choices []*Choice
}



type Test struct {
	Id string `json:"id"`
	UserID string `json:"user_id"`
	Question string `json:"question"`
}

