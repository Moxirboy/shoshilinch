package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
)



func FromreqToTeacherModel(
	req *dto.CreateTeacher,
) *models.User{
	return &models.User{
		FirstName: req.Firstname,
		LastName: req.Lastname,
		PhoneNumber: req.PhoneNumber,
		Password: req.Password,
		Role: "teacher",
	}
}
func FromTeacherModelToDto(
	teacher []*models.User,
) []*dto.Teacher{
	var res []*dto.Teacher
	for _,v:=range teacher{
		res=append(res,&dto.Teacher{
			ID: v.ID,
			Firstname: v.FirstName,
			Lastname: v.LastName,
			PhoneNumber: v.PhoneNumber,
			Password: v.Password,
			Role: "teacher",
		})
	}
	return res
}