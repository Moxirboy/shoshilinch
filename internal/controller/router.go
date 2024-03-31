package controller

import (
	"shoshilinch/internal/controller/v1/handler"
	"shoshilinch/internal/controller/v1/middleware"
	"shoshilinch/internal/service/usecase"
	Bot "shoshilinch/pkg/bot"
	"shoshilinch/pkg/log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(
	g *gin.Engine,
	log log.Logger,
	uc usecase.IUseCase,
	bot Bot.Bot,
) {
	NewInit(
		g.Group("/v1"),
		log,
		uc,
		bot,
	)
	NewFrontend(
		g.Group("/"),
		log,
	)
	url := ginSwagger.URL("swagger/doc.json")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func NewInit(
	g *gin.RouterGroup,
	log log.Logger,
	uc usecase.IUseCase,
	bot Bot.Bot,
) {
	handler.NewAdminHandler(
		g,
		uc.UserUseCase(),
		log,
	)
	handler.NewAuthHandler(
		g,
		uc.AuthUseCase(),
		uc.ClassUseCase(),
		uc.UserUseCase(),
		log,
	)
	handler.NewClassHandler(
		g,
		uc.ClassUseCase(),
		log,
	)
	handler.NewExamHandler(
		g,
		uc.ExamUseCase(),
		uc.ClassUseCase(),
		uc.UserUseCase(),
		bot,
		log,
	)
	handler.NewTestHandler(
		g,
		uc.TestUseCase(),
		bot,
		log,
	)
}

func NewFrontend(
	g *gin.RouterGroup,
	log log.Logger,
) {

	g.Static("/static", "internal/controller/templates")
	g.StaticFile("/", "internal/controller/templates/login.html")
	g.StaticFile("/sign", "internal/controller/templates/signUp.html")
	admin:=g.Group("/admin")
	admin.Use(middleware.AuthAdminMiddleware())
	admin.StaticFile("/", "internal/controller/templates/admin.html")
	{
		student:=g.Group("/student")
	{
		student.Use(middleware.AuthStudentMiddleware())
		student.StaticFile("/", "internal/controller/templates/student.html")
		student.StaticFile("/test", "internal/controller/templates/test.html")
	}
}
{
	teacher:=g.Group("/teacher")
	{
		teacher.Use(middleware.AuthTeacherMiddleware())
	teacher.StaticFile("/", "internal/controller/templates/teacher.html")
	teacher.StaticFile("/test/create", "internal/controller/templates/createTest.html")
	teacher.StaticFile("/class/create", "internal/controller/templates/classCreate.html")
	teacher.StaticFile("/exam/create", "internal/controller/templates/ExamCreate.html")
	teacher.StaticFile("/static/js/createClass.js", "internal/controller/templates/js/createClass.js")
	teacher.StaticFile("/exam/js/ExamCreate.js", "internal/controller/templates/js/ExamCreate.js")
}
}

}
