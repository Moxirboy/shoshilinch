package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
)

func FromReqToTestModelArray(
	req []dto.Tests,
	id string,
) (res []*models.Tests) {
for _,test:=range req{
	res = append(res, &models.Tests{
		Test: models.Test{
			Id:       test.Test.Id,
			Question: test.Test.Question,
			UserID:   id,
		},
		Choices: FromReqToChoiceModel(test.Choices),
	})
}
	return res
}

func FromReqToTestModel(
	req dto.Tests,
	id string,
) (res []*models.Tests) {

	res = append(res, &models.Tests{
		Test: models.Test{
			Id:       req.Test.Id,
			Question: req.Test.Question,
			UserID:   id,
		},
		Choices: FromReqToChoiceModel(req.Choices),
	})

	return res
}

func FromReqToChoiceModel(
	req []dto.Choice,
) (res []*models.Choice) {
	choices := &models.Choice{}
	for _, choice := range req {

		choices = &models.Choice{
			Id:         choice.Id,
			QuestionId: choice.QuestionId,
			Text:       choice.Text,
			IsTrue:     choice.IsTrue,
		}
		res = append(res, choices)
	}

	return res
}

func FromReqToAnswerModel(
	req dto.Answer,
) models.Answer {
	return models.Answer{
		QuestionId: req.QuestionId,
		TrueAnswer: req.TrueAnswer,
		Submitted:  req.Submitted,
	}
}
