package models

type Choice struct{
	Id string `json:"id"`
	QuestionId string `json:"question_id"`
	Text string `json:"text"`
	IsTrue bool `json:"is_true"`
}