package repo

import (
	"context"
	"shoshilinch/internal/models"
)



type ChoiceRepository interface{
	CreateChoice(
		ctx context.Context,
		test []*models.Tests,
	) error
}