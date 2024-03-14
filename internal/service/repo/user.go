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
		user models.User,
	) (
	 string,
	 error,
	)
}