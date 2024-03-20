package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
)


func FromReqToTestModel(
	req []*dto.Tests,
) (res []*models.Tests){

	for _,test:=range req {
		res=append(res,&models.Tests{
			Test:models.Test{
				Id: test.Test.Id,
				Question: test.Test.Question,
				UserID: test.Test.UserID,
			},
			Choices:FromReqToChoiceModel(test.Choices),
		})
	}
	return res
}



func FromReqToChoiceModel(
	req []*dto.Choice,
) (res []*models.Choice){
	choices:=&models.Choice{}
	for _,choice:=range req{
		choices=&models.Choice{
			Id: choice.Id,
			QuestionId: choice.QuestionId,
			Text: choice.Text,
			IsTrue: choices.IsTrue,
		}
		res=append(res, choices)
	}
	return res
}
