package repo
import (
	"context"
	"shoshilinch/internal/models"
)

type TestRepository interface{
	Create(
		ctx context.Context,
		test []*models.Tests,
	)error
}