package dto

type Tests struct {
	Test    Test
	Choices []Choice
}

type Test struct {
	Id       string `json:"id"`
	UserID   string `json:"user_id"`
	Question string `json:"question"`
}

type Answer struct {
	QuestionId string `json:"Question"`
	TrueAnswer string `json:"Answer"`
	Submitted  string `json:"submitted"`
}

type Score struct{
	Number string `json:"number"`
}