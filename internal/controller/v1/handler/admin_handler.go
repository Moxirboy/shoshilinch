package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/controller/v1/middleware"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-gonic/gin"
)

type adminHandler struct{
	uc usecase.UserUsecase
	log log.Logger
}
func NewAdminHandler(
	group *gin.RouterGroup,
	uc usecase.UserUsecase,
	log log.Logger,
) {
	handler:=&adminHandler{
		uc:uc,
		log:log,
	}
	g:=group.Group("/admin")
	g.Use(middleware.AuthMiddleware())
	g.POST("/create_teacher",handler.CreateTeacher)
	g.GET("/teachers",handler.GetTeachers)
}
// @Router /v1/ [ost]
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

func (h *adminHandler) GetTeachers(c *gin.Context){
	teachers,err:=h.uc.GetTeachers(c)
	if err!=nil{
		h.log.Info(
			"handler.admin.GetTeachers error: ",
			err,
		)
		utils.SendResponse(
			c,
			nil,
			err,
		)
	}
	
	res:=FromTeacherModelToDto(teachers)
	c.JSON(200,res)
}