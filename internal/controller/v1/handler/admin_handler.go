package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-gonic/gin"
)

type adminHandler struct{
	uc usecase.UserUsecase
	log log.Logger
}
func New(
	uc usecase.UserUsecase,
	log log.Logger,
) {

}

func (h *adminHandler) CreateTeacher(c *gin.Context){
	req:=dto.CreateTeacher{}

	if err:=c.ShouldBindJSON(&req);err!=nil{
		utils.SendResponse(c,nil,err)
	}
	teacher:=FromreqToTeacherModel(&req)
	id,err:=h.uc.CreateUser(
		c,
		teacher,
	)
	if err!=nil{
		h.log.Info(
			"handler.admin.CreateTeacher error: ",
			err,
		)
		utils.SendResponse(
			c,
			nil,
			err,
		)
	}
	utils.SendResponse(
		c,
		id,
		nil,
	)
}

