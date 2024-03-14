package server

import (
	"shoshilinch/internal/config"
	"shoshilinch/pkg/log"

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

func (s *server) Run() error{
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return(r.Run(":5005"))
}