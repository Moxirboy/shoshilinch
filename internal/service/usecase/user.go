package usecase

import (
	"context"
	
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type userUsecase struct{
	repo repo.UserRepository
	class repo.ClassRepository
	log log.Logger
}

func NewUserUsecase(
	repo repo.UserRepository,
	class repo.ClassRepository,
	log log.Logger,
) UserUsecase{
	return &userUsecase{
		repo: repo,
		class: class,
		log: log,
	}
}
func (uc *userUsecase) CleanUp(
	ctx context.Context,
)error {
	return uc.repo.CleanUp(ctx)
}
func (uc *userUsecase) CreateUser(ctx context.Context,
user *models.User,
) (
string,
error,
){
	if user.Role!=models.Admin{
	exist,err:=uc.repo.GetExist(ctx,user.PhoneNumber)
	if err!=nil{
		uc.log.Info("usecase.user.createUser error: ",err)
		return "",err
	}
	uc.log.Info("usecase.user.createUser exist: ",exist)
	if exist{
		return	"",nil
	}
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
	uc.log.Info("usecase.user.GetUser pass: ",pass)
	uc.log.Info("usecase.user.GetUser password: ",password)
	if password!=pass{
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
	string,
	string,
	error,
){
	exist,err:=uc.class.GetExist(
		ctx,
		className,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	if !exist {
		return "","","",models.ErrClassNotFound
	}
	pass,err:=uc.class.GetPassword(
		ctx,
		className,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	if pass!=user.Password{
		return "","","",models.ErrPasswordNotMatch
	}
	id,err:=uc.CreateUser(ctx,user)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	classId,err:=uc.class.GetId(
		ctx,
		className,
	)
	if err!=nil{
		
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	uc.log.Info(classId)
	err=uc.class.BindClassStudent(
		ctx,
		classId,
		id,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	teacherId,err:=uc.class.GetTeacherIdbyName(
		ctx,
		className,
	)
	if err!=nil{
		uc.log.Info("usecase.user.getexist error: ",err)
		return "","","",err
	}
	return id,classId,teacherId,nil
}

func (uc *userUsecase) GetTeachers(
	ctx context.Context,
)(
	[]*models.User,
	error,
){
	teachers,err:=uc.repo.GetTeachers(ctx)
	if err!=nil{
		uc.log.Info("usecase.user.getTeachers error: ",err)
		return nil,err
	}
	return teachers,nil
}

func (uc *userUsecase) GetUserInfo(
	ctx context.Context,
	id string,
)(*models.User,error){
	return uc.repo.GetUserInfo(
		ctx,
		id,
	)
}