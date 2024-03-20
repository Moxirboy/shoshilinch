package handler

import (
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type examHandler struct{
	uc usecase.ExamUsecase
	log log.Logger
}

func NewExamHandler(
	uc usecase.ExamUsecase,
	log log.Logger,
){

}

func (h *examHandler) GetExist(c *gin.Context){
	s:=sessions.Default(c)
	id:=s.Get("ClassId").(string)
	err:=h.uc.GetExist(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
	}
	utils.SendResponse(c,nil,nil)
}

func (h *examHandler) CreateExam(c *gin.Context){
	s:=sessions.Default(c)

	id:=s.Get("ID").(string)
	examId,err:=h.uc.Create(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
	}
	utils.SendResponse(c,examId,nil)
}