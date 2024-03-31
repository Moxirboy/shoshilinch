package models


type Tests struct{
	Test Test
	Choices []*Choice
}




type Testss struct {
    Test    Test       `json:"test"`
    Choices []Choices `json:"choices"`
}

type Test struct {
    Id       string `json:"id"`
    UserID   string `json:"user_id"`
    Question string `json:"question"`
}

type Choices struct {
    Id         string `json:"id"`
    QuestionId string `json:"question_id"`
    Text       string `json:"text"`
    IsTrue     bool   `json:"is_true"`
}


type Answer struct{
    Id string
	QuestionId string `json:"Question"`
	TrueAnswer string `json:"Answer"`
	Submitted  string `json:"submitted"`
}