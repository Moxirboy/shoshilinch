package usecase

import (
	"context"
	"shoshilinch/internal/config"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/jwt"
	"shoshilinch/pkg/log"
	
)


type authUseCase struct{
	repo repo.AuthRepository
	log log.Logger
}
func NewAuthUseCase(
	repo repo.AuthRepository,
	log log.Logger,
) AuthUseCase {
	return &authUseCase{
		repo: repo,
		log: log,
	}
}
func (uc authUseCase) New(
	ctx context.Context,
	id,role string,
)(
	string,
	error,
){
	return uc.generate(id,role)
}

func (uc authUseCase) Check(
	ctx context.Context,
	token string,
)(id ,role string,err error){
	cfg:=config.Load()
	data,err:=jwt.ParseJWT(token,cfg.JWT.SecretKey)
	if err!=nil{
		uc.log.Info("usecase.auth.Check error: ",err)
		return "","",err
	}
	return data.UserID,data.Role,nil
}


func (uc authUseCase) generate(
	id, role string,
)(string,error){
	cfg:=config.Load()
	token,err:=jwt.GenerateJWT(id,role,cfg.JWT.SecretKey)
	if err!=nil{
		uc.log.Info("usecase.auth.generate error: ",err)
		return "",err
	}

	return token,nil
}