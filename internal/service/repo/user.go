package repo

import (
	"context"
	"shoshilinch/internal/models"
)


type UserRepository interface {
	GetExist(
		ctx context.Context,
		phoneNumber string,
		) (
		bool,
		error,
	)
	Create(
		ctx context.Context,
		user *models.User,
	) (
	 string,
	 error,
	)
	Get(
		ctx context.Context,
		phoneNumber string,
	) (string,string,string,error)
	CleanUp(
		ctx context.Context,
	) error
	GetTeachers(
		ctx context.Context,
	) ([]*models.User,error)
	GetUserInfo(
		ctx context.Context,
		id string,
	)(*models.User,error)
}