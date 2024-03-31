package usecase

import (
	"context"
	"errors"

	"shoshilinch/internal/models"
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
	userId string,
	id string,
)(string,bool, error){
	exist,err :=uc.repo.GetExist(
		ctx,
		id,
	)
	if err!=nil{
		uc.log.Info(err)
		return "",false,err
	}
	if !exist{
		return "",false,errors.New("No")
	}
	ids,err:=uc.repo.GetId(
		ctx,
		id,
	)
	if err!=nil{
		uc.log.Info(err)
		return "",false,err
	}
	attempted,err:=uc.repo.GetAttempted(
		ctx,
		userId,
		ids,
	)
	if err!=nil{
		uc.log.Info(err)
		return "",false,err
	}
	if attempted {
		return "",false, errors.New("attempted")
	}
	
	return ids,true,nil
}
func (uc *examUsecase)MakeAttempted(
	ctx context.Context,
	id string,
	examId string,
)error{
	
	return uc.repo.MakeAttempted(ctx,id,examId)
}
func (uc *examUsecase) UpdateScore(
	ctx context.Context,
	at models.Attempted,
) error {
	return uc.repo.UpdateScore(
		ctx,
		at,
	)
}

func (uc *examUsecase) GetScore(
	ctx context.Context,
	id string,
) (int ,error){
	return uc.repo.GetScore(
		ctx,
		id,
	)
}
func (uc *examUsecase)GetStudents(
	ctx context.Context,
	id string,
)([]*models.Students,error){
	return uc.repo.GetStudents(
		ctx,
		id,
	)
}