package handler

import (
	"fmt"
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/controller/v1/middleware"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/usecase"
	Bot "shoshilinch/pkg/bot"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"
	"strconv"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type examHandler struct{
	uc usecase.ExamUsecase
	class usecase.ClassUsecase
	user usecase.UserUsecase
	bot Bot.Bot
	log log.Logger
}

func NewExamHandler(
	group *gin.RouterGroup,
	uc usecase.ExamUsecase,
	class usecase.ClassUsecase,
	user usecase.UserUsecase,
	bot Bot.Bot,
	log log.Logger,
){
	handler:=&examHandler{
		uc:uc,
		class:class,
		user:user,
		bot: bot,
		log:log,
	}
	g:=group.Group("/exam")
	g.Use(middleware.AuthMiddleware())

	g.POST("/create",handler.CreateExam)
	g.GET("/get_exist",handler.GetExist)
	g.GET("/attempt",handler.Attempted)
	g.POST("/score",handler.CreateScore)
	g.GET("/students",handler.GetScores)
}

func (h *examHandler) GetExist(c *gin.Context){
	s:=sessions.Default(c)
	var id =s.Get("classId").(string)
	var userId=s.Get("ID").(string)

	examId,exist,err:=h.uc.GetExist(
		c.Request.Context(),
		userId,
		id,
	)
	
	if err!=nil{
		if err.Error()=="No"{
			c.JSON(
				200,
				gin.H{
					"message":"NO",
				},
			)
			return
		} 
		if err.Error()=="attempted"{

			score,err:=h.uc.GetScore(
				c.Request.Context(),
				userId,
			)
			if err!=nil{
				c.JSON(200,
					gin.H{
						"exist":false,
					})	
					return
			}
			c.JSON(
				200,
				gin.H{
					"message":"attempted",
					"score": score,
				},
			)
			return
		}else {
			c.JSON(200,
				gin.H{
					"exist":false,
				})	
				return
		}
		
			}

	s.Set("examId",examId)
	s.Save()
	c.JSON(200,
	gin.H{
		"exist":exist,
	})
}

func (h *examHandler) CreateExam(c *gin.Context){
	req:=dto.Exam{}
	if err:=c.ShouldBind(&req);err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}

	id,err:=h.class.GetClass(c.Request.Context(),req.ClassName)
	if err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}

	examId,err:=h.uc.Create(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}
	utils.SendResponse(c,examId,nil)
}

func (h *examHandler) Attempted(c *gin.Context){
	s:=sessions.Default(c)
	var id= s.Get("ID").(string)
	var examid =s.Get("examId").(string)
	err:=h.uc.MakeAttempted(
		c.Request.Context(),
		id,
		examid,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
		return
	}
	c.JSON(200,gin.H{
		"status":"200",
	})
}


func (h *examHandler) CreateScore(c *gin.Context){
	s:=sessions.Default(c)
	req:=dto.Score{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		h.log.Info(err)
		utils.SendResponse(c,nil,err)
		return
	}
	id:=s.Get("ID").(string)
	teacherId:=s.Get("TID").(string)
	examId:=s.Get("examId").(string)
	h.log.Info(id)
	info,err:=h.user.GetUserInfo(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		h.log.Info(err)
		utils.SendResponse(c,nil,err)
		return
	}
	score,err:=strconv.Atoi(req.Number)
	if err!=nil{
		h.log.Info(err)
		utils.SendResponse(c,nil,err)
		return
	}
	h.uc.UpdateScore(
		c.Request.Context(),
		models.Attempted{
			Score: score,
			UserId: id,
			TeacherId: teacherId,
			ExamId: examId,
		},
	)


	message:=fmt.Sprintf(
		"Ism:%s\nFamiliya:%s\nTo'g'ri javoblar soni:%s\n",
		info.FirstName,
		info.LastName,
		req.Number,
	)
	h.log.Info(message)
	h.bot.SendNotification(
		message,
	)
	c.JSON(200,gin.H{})
}


func (h *examHandler) GetScores(c *gin.Context){
	s:=sessions.Default(c)
	id:=s.Get("ID").(string)
	students,err:=h.uc.GetStudents(
		c.Request.Context(),
		id,
	)
	if err!=nil{
		h.log.Info(err)
		utils.SendResponse(c,nil,err)
		return
	}
	c.JSON(200,students)
} 