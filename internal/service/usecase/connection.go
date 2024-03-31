package usecase

import (
	"github.com/jmoiron/sqlx"
	"shoshilinch/internal/config"
	"shoshilinch/internal/service/repo/postgres"
	"shoshilinch/pkg/log"
)

type IUseCase interface {
	AuthUseCase() AuthUseCase
	ClassUseCase() ClassUsecase
	ExamUseCase() ExamUsecase
	UserUseCase() UserUsecase
	TestUseCase() TestUsecase
}
type UseCase struct {
	connections map[string]interface{}
}

const (
	_authUseCase  = "auth_use_case"
	_classUseCase = "class_use_case"
	_examUseCase  = "exam_use_case"
	_userUseCase  = "user_use_case"
	_testUseCase  = "test_use_case"
)

func New(
	cfg *config.Config,
	db *sqlx.DB,
	log log.Logger,
) IUseCase {
	var connections = make(map[string]interface{})
	connections[_authUseCase] = NewAuthUseCase(
		postgres.NewAuthRepository(
			db,
			log,
		),
		log,
	)
	connections[_classUseCase] = NewClassUsecase(
		postgres.NewClassRepository(
			db,
			log,
		), log,
	)
	connections[_examUseCase] = NewExamUsecase(
		postgres.NewExamRepository(
			db,
			log,
		), log,
	)
	connections[_userUseCase] = NewUserUsecase(
		postgres.NewUserRepository(
			db,
			log,
		),
		postgres.NewClassRepository(
			db,
			log,
		),
		log,
	)
	connections[_testUseCase]=NewTestUseCase(
		postgres.NewTestRepository(
			db,
			log,
		), 
		postgres.NewChoiceRepository(
			db,
			log,
		),
		log,
	)
	return &UseCase{
		connections: connections,
	}
}

func (c *UseCase) AuthUseCase() AuthUseCase {
	return c.connections[_authUseCase].(AuthUseCase)
}
func (c *UseCase) ClassUseCase() ClassUsecase {
	return c.connections[_classUseCase].(ClassUsecase)
}
func (c *UseCase) ExamUseCase() ExamUsecase {
	return c.connections[_examUseCase].(ExamUsecase)
}
func (c *UseCase) UserUseCase() UserUsecase {
	return c.connections[_userUseCase].(UserUsecase)
}
func (c *UseCase) TestUseCase() TestUsecase {
	return c.connections[_testUseCase].(TestUsecase)
}
