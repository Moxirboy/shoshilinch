package handler

import (
	"errors"
	"shoshilinch/internal/controller/v1/dto"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/usecase"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	uca   usecase.AuthUseCase
	class usecase.ClassUsecase
	uc    usecase.UserUsecase
	log   log.Logger
}

func NewAuthHandler(
	group *gin.RouterGroup,
	uca usecase.AuthUseCase,
	class usecase.ClassUsecase,
	uc usecase.UserUsecase,
	log log.Logger,
) {
	auth := authHandler{
		uca,
		class,
		uc,
		log,
	}
	g := group.Group("/auth")
	{
		g.POST("/login", auth.Login)
		g.POST("/sign-up", auth.SignUp)
	}
}

// @Router /v1/auth/Login [post]
func (h *authHandler) Login(c *gin.Context) {
	s := sessions.Default(c)
	req := dto.Login{}
	if err := c.ShouldBind(
		&req,
	); err != nil {
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		c.JSON(404, dto.InvalidParams{
			Name:   "invalid params",
			Reason: "coudn`t bind",
		})
	}
	invalidParams := utils.Validate(req)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}
	id, role, err := h.uc.GetUser(c.Request.Context(), req.PhoneNumber, req.Password)

	if errors.Is(err, models.ErrUserNotFound) {
		utils.SendResponse(c, nil, models.ErrUserNotFound)
		return
	} else if errors.Is(err, models.ErrPasswordNotMatch) {
		utils.SendResponse(c, nil, models.ErrPasswordNotMatch)
		return
	} else if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	token, err := h.uca.New(c, id, role)
	if err != nil {
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		utils.SendResponse(c, nil, err)
		return
	}
	s.Set("ID", id)
	s.Set("Role", role)
	if role == "student" {
		class, err := h.class.Get(c.Request.Context(), id)
		if err != nil {
			h.log.Info(
				"handler.auth.login error: ",
				err,
			)
			utils.SendResponse(c, nil, err)
			return
		}
		s.Set("classId", class)
		TID,err:=h.class.GetTeacherIdbyId(c.Request.Context(),class)
		if err != nil {
			h.log.Info(
				"handler.auth.login error: ",
				err,
			)
			utils.SendResponse(c, nil, err)
			return
		}
		s.Set("TID",TID)
	}
	s.Save()
	
	c.JSON(200, dto.AuthResponse{
		Token: token,
		Role:  role,
	})
}

// @Summary 	Sign-Up user by email
// @Description This api can Sign-Up new user by phone
// @Tags 		Sign-Up
// @Accept 		json
// @Produce 	json
// @Param body  body dto.SignUp true "Sign"
// @Failure 400 string Error response
// @Router /v1/auth/sign-up [post]
func (h *authHandler) SignUp(c *gin.Context) {
	s := sessions.Default(c)
	req := dto.SignUp{}

	if err := c.ShouldBind(&req); err != nil {
		h.log.Info(
			"handler.auth.login error: ",
			err,
		)
		c.JSON(200, dto.InvalidParams{
			Name:   "invalid params",
			Reason: "could`t bind",
		})
		return
	}
	h.log.Info(req)
	id, classId, TID, err := h.uc.SignUp(
		c.Request.Context(),
		ReqToModel(&req),
		req.ClassName,
	)
	if err != nil {
		h.log.Info(
			"handler.auth.signUp error: ",
			err,
		)
		c.JSON(200, dto.InvalidParams{
			Name:   "invalid params",
			Reason: err.Error(),
		})		
		return
	}
	token, err := h.uca.New(
		c.Request.Context(),
		id,
		"student",
	)
	if err != nil {
		h.log.Info(
			"handler.auth.signUp error: ",
			err,
		)
		c.JSON(200, dto.InvalidParams{
			Name:   "invalid params",
			Reason: err.Error(),
		})		
		return
	}
	s.Set("ID", id)
	s.Set("Role", "student")
	s.Set("classId", classId)
	s.Set("TID",TID)
	s.Save()
	
	data := dto.AuthResponse{
		Token: token,
		Role:  "student",
	}
	c.JSON(
		200,
		data,
	)
}
