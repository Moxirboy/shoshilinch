package repo

import (
	"context"
	"shoshilinch/internal/models"
)



type ChoiceRepository interface{
	Create(
		ctx context.Context,
		test []*models.Tests,
	)error
}