package usecase

import (
	"context"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type classUsecase struct {
	repo repo.ClassRepository
	log log.Logger
}


func NewClassUsecase(
	repo repo.ClassRepository,
	log log.Logger,
) ClassUsecase{
	return &classUsecase{
		repo: repo,
		log: log,
	}
}

func (uc *classUsecase) Create(
	ctx context.Context,
	class *models.Class,
) error{
	return uc.repo.Create(ctx,class)
}

func (uc *classUsecase) GetAll(
	ctx context.Context,
	id string,
) ([]*models.Class,error){
	return uc.repo.GetAll(
		ctx,
		id,
	)
}

func (uc *classUsecase) Get(
	ctx context.Context,
	id string,
) (string,error){
	return uc.repo.Get(
		ctx,
		id,
	)
}
func (uc *classUsecase) GetClass(
	ctx context.Context,
	name string,
) (string,error){
	return uc.repo.GetId(
		ctx,
		name,
	)
}

func (uc *classUsecase) GetTeacherIdbyId(
	ctx context.Context,
	id string,
) (string,error){
	return uc.repo.GetTeacherIdbyId(
		ctx,
		id,
	)
}