package usecase

import (
	"context"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type testUsecase struct {
	repo       repo.TestRepository
	choiceRepo repo.ChoiceRepository
	log        log.Logger
}

func NewTestUseCase(
	repo repo.TestRepository,
	choiceRepo repo.ChoiceRepository,
	log log.Logger,
) TestUsecase {
	return &testUsecase{
		repo: repo,
		log:  log,
	}
}

func (uc *testUsecase) CreateTest(
	ctx context.Context,
	test []*models.Tests,
) error {
	err := uc.repo.Create(
		ctx,
		test,
	)
	if err != nil {
		uc.log.Info(err)
		return err
	}

	err = uc.repo.CreateChoice(
		ctx,
		test,
	)

	if err != nil {
		uc.log.Info(err)
		return err
	}

	return nil
}
func (uc *testUsecase) GetRandom(
	ctx context.Context,
	id string,
) (*models.Testss, error) {
	return uc.repo.GetRandom(ctx, id)
}

func (uc *testUsecase) SaveAnswer(
	ctx context.Context,
	id string,
	answer models.Answer,
) error {
	return uc.repo.SaveAnswer(
		ctx,
		id,
		answer,
	)
}
