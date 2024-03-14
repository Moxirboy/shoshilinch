package repo

import (
	"context"
	"time"
	"shoshilinch/internal/models"
)


type AuthRepository interface {

	Create(
		ctx context.Context,
		token string,
		id string,
		role string,
		date time.Time,
	) error

	GetExist(
		ctx context.Context,
		token string,
		date time.Time,
	)(
		bool,
		error,
	)
	Get(
		ctx context.Context,
		token string,
		date time.Time,
	) (
		*models.Token,
		error,
	)
	CleanUp(
		ctx context.Context,
		
	)error
}