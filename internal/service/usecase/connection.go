package usecase

import (
	"database/sql"
	"shoshilinch/internal/config"
	"shoshilinch/internal/service/repo/postgres"
	"shoshilinch/pkg/log"
)


type IUseCase interface {
	AuthUseCase() AuthUseCase
	
}
type UseCase struct{
	connection map[string]interface{}
}
 const (
	_authUseCase="auth_use_case"
	_classUseCase="class_use_case"
	_examUseCase="exam_use_case"
	_userUseCase="user_use_case"
 )

 func New(
	cfg *config.Config,
	db *sql.DB,
	log log.Logger,
 ) *UseCase{
	var connections =make(map[string]interface{})
	connections[_authUseCase]=NewAuthUseCase(
		postgres.NewAuthRepository(
			db,
			log,
		),
		log,
	)
	connections[_classUseCase]=NewClassUsecase(
		postgres.NewClassRepository(
			db,
			log,
		),log,
	)
	connections[_examUseCase]=NewExamUsecase(
		postgres.NewExamRepository(
			db,
			log,
		),log,
	)
	connections[_userUseCase]=NewUserUsecase(
		postgres.NewUserRepository(
			db,
			log,
		),log,
	)
	return &UseCase{
		connection: connections,
	}
 }
 