package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
)


func ReqToModel(
	req *dto.SignUp,
) *models.User{
	return &models.User{
		FirstName: req.Firstname,
		LastName: req.Lastname,
		PhoneNumber: req.PhoneNumber,
		Password: req.Password,
		Role: "student",
	}
}