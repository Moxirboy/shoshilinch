package handler

import (
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-gonic/gin"
)


type testHandler struct {
	uc usecase.TestUsecase
	log log.Logger
}

func NewTestHandler(
	uc usecase.TestUsecase,
	log log.Logger,
) {

}

func (h *testHandler) CreateTest(c *gin.Context){
	req:=[]*dto.Tests{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		utils.SendResponse(c,nil,err)
	}
	Tests:=FromReqToTestModel(req)
	err:=h.uc.CreateTest(
		c.Request.Context(),
		Tests,
	)
	if err!=nil{
		utils.SendResponse(c,nil,err)
	}
	utils.SendResponse(c,Tests,nil)
}


func (h *testHandler) GetTests(c *gin.Context){
	
}