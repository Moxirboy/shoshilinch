package handler

import (
	"errors"
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-gonic/gin"
)



type authHandler struct{
	uca usecase.AuthUseCase
	uc usecase.UserUsecase
	log log.Logger
}

func NewAuthHandler(
	group *gin.RouterGroup,
	uca usecase.AuthUseCase,
	uc usecase.UserUsecase,
	log log.Logger,
){
	auth := authHandler{
		uca,
		uc,
		log,
	}
	_=auth
}

func (h *authHandler) Login(c *gin.Context){
	req:=dto.Login{}
	if err:=c.ShouldBind(
		&req,
	);err!=nil{
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		c.JSON(404,dto.InvalidParams{
			Name:"invalid params",
			Reason: "coudn`t bind",
		},)
	}
	invalidParams:=utils.Validate(req)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
	}
	id,role,err:=h.uc.GetUser(c.Request.Context(),req.PhoneNumber,req.Password)
	if errors.Is(err,models.ErrUserNotFound){
		utils.SendResponse(c,nil,models.ErrUserNotFound)
	}else if (errors.Is(err,models.ErrPasswordNotMatch)){
		utils.SendResponse(c,nil,models.ErrPasswordNotMatch)
	}else{
		utils.SendResponse(c,nil,err)
	}
	token,err:=h.uca.New(c,id,role)
	if err!=nil{
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		utils.SendResponse(c,nil,err)
	}
	utils.SendResponse(c,dto.AuthResponse{
		Token:token,
		Role:role,
	},nil)
}
// @Router /v1/sign-up [post]
func (h *authHandler) SignUp(c *gin.Context){
	req:=dto.SignUp{}

	if err:=c.ShouldBind(&req);err!=nil{
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		c.JSON(404,dto.InvalidParams{
			Name:"invalid params",
			Reason: "coudn`t bind",
		},)
	}
	invalidParams:=utils.Validate(req)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
	}

	id,err:=h.uc.SignUp(
		c.Request.Context(),
		ReqToModel(&req),
		req.ClassName,
	)
	if err!=nil{
		h.log.Info(
			"handler.auth.signUp error: ",
			err,
		)
		utils.SendResponse(c,nil,err)
	}
	token,err:=h.uca.New(
		c.Request.Context(),
		id,
		"student",
	)
	utils.SendResponse(c,dto.AuthResponse{
		Token:token,
		Role:"student",
	},nil)
}
