package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type classHandler struct{
	uc usecase.ClassUsecase
	log log.Logger
}

func NewClassHandler(
	uc usecase.ClassUsecase,
	log log.Logger,
){

}


func (h *classHandler) CreateClass(c *gin.Context){
	s :=sessions.Default(c)
	req:=dto.Class{}
	if err:=c.ShouldBind(&req);err!=nil{
		utils.SendResponse(c,nil,err)
	}
	id:=s.Get("ID").(string)
	class:=FromReqToClassModel(&req,id)
	err:=h.uc.Create(c.Request.Context(),class)
	if err!=nil{
		utils.SendResponse(c, nil,err)
	}
	utils.SendResponse(c,nil,nil)
}


func (h *classHandler) GetAllClass(c *gin.Context){
	s:=sessions.Default(c)
	id:=s.Get("ID").(string)
	classes,err:=h.uc.GetAll(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
	}
	utils.SendResponse(c,FromClassesToRes(classes),nil)
	
}