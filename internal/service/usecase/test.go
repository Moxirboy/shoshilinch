package usecase

import (
	"context"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type testUsecase struct{
	repo repo.TestRepository
	choiceRepo repo.ChoiceRepository
	log log.Logger
}

func NewTestUseCase(
	repo repo.TestRepository,
	choiceRepo repo.ChoiceRepository,
	log log.Logger,
)TestUsecase{
	return &testUsecase{
		repo: repo,
		log: log,
	}
}

func (uc *testUsecase) CreateTest(
	ctx context.Context,
	test []*models.Tests,
) error{
	
	err:=uc.repo.Create(
		ctx,
		test,
	)
	if err!=nil{
		return err
	}

	err=uc.choiceRepo.Create(
		ctx,
		test,
	)

	if err!=nil{
		return err
	}

	return nil
}