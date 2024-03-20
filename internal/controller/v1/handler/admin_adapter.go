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
		Password: req.PhoneNumber,
		Role: "teacher",
	}
}