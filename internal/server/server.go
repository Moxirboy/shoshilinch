package server

import (
	"context"
	"fmt"
	_ "shoshilinch/cmd/docs" //
	"shoshilinch/internal/config"
	"shoshilinch/internal/controller"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/usecase"
	Bot "shoshilinch/pkg/bot"
	"shoshilinch/pkg/log"
	"shoshilinch/pkg/postgres"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type server struct {
	cfg *config.Config
	log log.Logger
}

func New(
	cfg *config.Config,
	log log.Logger,
) *server {
	return &server{
		cfg: cfg,
		log: log,
	}
}

func (s *server) Run() error {
	db, err := postgres.DB(&s.cfg.Postgres)
	if err != nil {
		fmt.Println(err)
	}
	uc := usecase.New(
		s.cfg,
		db,
		s.log,
	)
	context := context.Background()
	err = uc.UserUseCase().CleanUp(context)
	if err != nil {
		s.log.Fatal("Cannot create Admin error: ", err)
		return err
	}
	id, err := uc.UserUseCase().CreateUser(
		context,
		&models.User{
			FirstName:   s.cfg.ADMIN.Firstname,
			LastName:    s.cfg.ADMIN.Lastname,
			Role:        models.Admin,
			Password:    s.cfg.ADMIN.Password,
			PhoneNumber: s.cfg.ADMIN.PhoneNumber,
		},
	)
	s.log.Info("Admin created with id: ", id)
	if err != nil {
		s.log.Fatal("Cannot create Admin error: ", err)
		return err
	}
	bot, err := Bot.NewBotConfig(s.cfg)
	if err != nil {
		s.log.Fatal("Cannot initialize bot error: ", err)
		return err
	}
	NewBOT := Bot.NewBot(bot)
	s.log.Info(NewBOT)
	r := gin.Default()
	store := cookie.NewStore([]byte("nimadur")) // Change "secret" to your desired session secret
	r.Use(sessions.Sessions("mysession", store))
	controller.New(r, s.log, uc, NewBOT)

	return (r.Run(":5005"))
}
