package repo

import (
	"context"
	"shoshilinch/internal/models"
)

type TestRepository interface {
	Create(
		ctx context.Context,
		test []*models.Tests,
	) error
	GetRandom(
		ctx context.Context,
		id string,
	) (*models.Testss, error)
	CreateChoice(
		ctx context.Context,
		test []*models.Tests,
	) error
	SaveAnswer(
		ctx context.Context,
		id string,
		answer models.Answer,
	) error
}
