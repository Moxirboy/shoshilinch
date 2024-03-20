package repo
import (
	"context"
	"shoshilinch/internal/models"
)

type ClassRepository interface{
	GetExist(
		ctx context.Context,
		name string,
	) (bool,error)
	GetPassword(
		ctx context.Context,
		name string,
	) (string,error)
	Create(
		ctx context.Context,
		class *models.Class,
	)error
	GetAll(
		ctx context.Context,
		id string,
	) ([]*models.Class,error)
}