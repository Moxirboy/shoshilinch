package usecase

import (
	"context"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)



type examUsecase struct{
	repo repo.ExamRepository
	log log.Logger
}

func NewExamUsecase(
	repo repo.ExamRepository,
	log log.Logger,
) ExamUsecase {
 return &examUsecase{
	repo:repo,
	log:log,
}
}

func (uc *examUsecase) Create(
	ctx context.Context,
	id string,
)( string, error){
	return uc.repo.Create(ctx,id)
}

func (uc *examUsecase) GetExist(
	ctx context.Context,
	id string,
) error{
	return uc.repo.GetExist(
		ctx,
		id,
	)
}