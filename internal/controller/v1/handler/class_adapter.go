package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
)



func FromReqToClassModel(
	req *dto.Class,
	id string,
) *models.Class{
	return &models.Class{
		Name: req.Name,
		Password: req.Password,
		UserId: id,
	}
}

func FromClassesToRes(
	class []*models.Class,
) []*dto.Class{
	res:=[]*dto.Class{}
	for _,c:=range class{
		res=append(res,FromClassToRes(c))
	}
	return res
}

func FromClassToRes(
	res	*models.Class,
) *dto.Class{
	return &dto.Class{
		ID: res.ID,
		Name:res.Name,
		Password: res.Password,
	}
}