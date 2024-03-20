package usecase

import (
	"context"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
	"shoshilinch/internal/models"
)

type userUsecase struct{
	repo repo.UserRepository
	class repo.ClassRepository
	log log.Logger
}

func NewUserUsecase(
	repo repo.UserRepository,
	log log.Logger,
) UserUsecase{
	return &userUsecase{
		repo: repo,
		log: log,
	}
}

func (uc *userUsecase) CreateUser(ctx context.Context,
user *models.User,
) (
string,
error,
){
	exist,err:=uc.repo.GetExist(ctx,user.PhoneNumber)
	if err!=nil{
		uc.log.Info("usecase.user.createUser error: ",err)
		return "",err
	}
	if !exist{
		return	"",nil
	}
	id,err:=uc.repo.Create(ctx,user)
	if err!=nil{
		uc.log.Info("usecase.user.createUser error: ",err)
		return "",err
	}
	return id,err
}

func (uc *userUsecase) GetUser(
	ctx context.Context,
	phoneNumber string,
	password string,
)(string,string,error){
	exist,err:=uc.repo.GetExist(
		ctx,
		phoneNumber,
	)
	if err!=nil{
		uc.log.Info("usecase.user.GetUser error: ",err)
		return "","",err
	}
	if !exist{
		return "","",nil
	}
	id,role,pass,err:=uc.repo.Get(
		ctx,
		phoneNumber,
	)
	if err!=nil{
		uc.log.Info("usecase.user.GetUser error: ",err)
		return "","",err
	}
	if password==pass{
		return "","",models.ErrPasswordNotMatch
	}
	return id,role,nil
}

func (uc *userUsecase) SignUp(
	ctx context.Context,
	user *models.User,
	className string,
)(
	string,
	error,
){
	exist,err:=uc.class.GetExist(
		ctx,
		className,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "",err
	}
	if !exist {
		return "",models.ErrClassNotFound
	}
	pass,err:=uc.class.GetPassword(
		ctx,
		className,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "",err
	}
	if pass!=user.Password{
		return "",models.ErrPasswordNotMatch
	}
	return uc.CreateUser(ctx,user)
}