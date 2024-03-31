package handler

import (
	"fmt"
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/controller/v1/middleware"
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
	group *gin.RouterGroup,
	uc usecase.ClassUsecase,
	log log.Logger,
){
	g:=group.Group("/class")
	handler:=&classHandler{
		uc:uc,
		log:log,
	}
	g.Use(middleware.AuthMiddleware())

	g.POST("/create",handler.CreateClass)
	g.GET("/get_all",handler.GetAllClass)
}


func (h *classHandler) CreateClass(c *gin.Context){
	s :=sessions.Default(c)
	req:=dto.Class{}
	if err:=c.ShouldBind(&req);err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}
	var id =s.Get("ID").(string)
	
	class:=FromReqToClassModel(&req,id)
	err:=h.uc.Create(c.Request.Context(),class)
	if err!=nil{
		utils.SendResponse(c, nil,err)
		return
	}
	utils.SendResponse(c,nil,nil)
}


func (h *classHandler) GetAllClass(c *gin.Context){
	s:=sessions.Default(c)
	var id =s.Get("ID").(string)
	classes,err:=h.uc.GetAll(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}
	fmt.Println(classes)
	res:=FromClassesToRes(classes)
	fmt.Println(&res)
	utils.SendResponse(c,res,nil)
	
}