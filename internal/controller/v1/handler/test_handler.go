package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/controller/v1/middleware"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/usecase"
	Bot "shoshilinch/pkg/bot"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
)

type testHandler struct {
	uc  usecase.TestUsecase
	log log.Logger
	bot Bot.Bot
}

func NewTestHandler(
	group *gin.RouterGroup,
	uc usecase.TestUsecase,
	bot Bot.Bot,
	log log.Logger,
) {
	handler := &testHandler{
		uc:  uc,
		log: log,
		bot: bot,
	}
	g := group.Group("/test")
	g.Use(middleware.AuthMiddleware())
	g.POST("/create", handler.CreateTest)
	g.GET("/get_all", handler.GetTests)
	g.POST("/save", handler.SaveAnswer)
	g.POST("/array",handler.CreateTestArray)
}

func (h *testHandler) CreateTest(c *gin.Context) {
	s := sessions.Default(c)
	req := dto.Tests{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Info(err)
		utils.SendResponse(c, nil, err)
		return
	}

	var id = s.Get("ID").(string)
	h.bot.SendNotification(h.log.LogArgsToString(req))
	Tests := FromReqToTestModel(req, id)
	err := h.uc.CreateTest(
		c.Request.Context(),
		Tests,
	)
	if err != nil {
		h.log.Info(err)
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Add other headers here to support OPTIONS preflighted requests, if necessary
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *testHandler) GetTests(c *gin.Context) {
	enableCors(c.Writer) // Call enableCors function passing the writer from gin.Context
	s := sessions.Default(c)
	var id =s.Get("TID").(string)

	tests, err := h.uc.GetRandom(
		c.Request.Context(),
		id,
	)
	if err != nil {
		h.log.Info("coming")
		h.log.Info("err: ", err)
		utils.SendResponse(c, nil, err)
		return
	}
	if tests == nil {
		tests := models.Testss{
			Test: models.Test{Id: "1", UserID: "user1", Question: "What is Go?"},
			Choices: []models.Choices{
				{Id: "c1", QuestionId: "1", Text: "A programming language", IsTrue: true},
				{Id: "c2", QuestionId: "1", Text: "A game", IsTrue: false},
			},
		}

		c.JSON(200, tests)
		return
	}
	h.log.Info(tests)
	c.JSON(200, tests)
}

func (h *testHandler) SaveAnswer(c *gin.Context) {
	s := sessions.Default(c)
	id := s.Get("ID").(string)
	
	req := dto.Answer{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	answer := FromReqToAnswerModel(req)
	err := h.uc.SaveAnswer(
		c.Request.Context(),
		id,
		answer,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	c.JSON(200,
		gin.H{})
}




func (h *testHandler) CreateTestArray(c *gin.Context) {
	s := sessions.Default(c)
	req := []dto.Tests{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Info(err)
		utils.SendResponse(c, nil, err)
		return
	}

	var id = s.Get("ID").(string)
	h.log.Info(req)
	//h.bot.SendNotification(h.log.LogArgsToString(req))
	Tests := FromReqToTestModelArray(req, id)
	err := h.uc.CreateTest(
		c.Request.Context(),
		Tests,
	)
	if err != nil {
		h.log.Info(err)
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, nil, nil)
}
