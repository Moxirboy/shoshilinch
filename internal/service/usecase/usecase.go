package usecase

import (
	"context"
	"shoshilinch/internal/models"
)

type AuthUseCase interface {
	New(
		ctx context.Context,
		id, role string,
	) (
		string,
		error,
	)
	Check(
		ctx context.Context,
		token string,
	) (id, role string, err error)
}

type UserUsecase interface {
	CreateUser(ctx context.Context,
		user *models.User,
	) (
		string,
		error,
	)
	GetUser(
		ctx context.Context,
		phoneNumber string,
		password string,
	) (string, string, error)
	SignUp(
		ctx context.Context,
		user *models.User,
		className string,
	) (
		string,
		string,
		string,
		error,
	)
	CleanUp(
		ctx context.Context,
	) error
	GetTeachers(
		ctx context.Context,
	) (
		[]*models.User,
		error,
	)
	GetUserInfo(
		ctx context.Context,
		id string,
	) (*models.User, error)
}

type ClassUsecase interface {
	Create(
		ctx context.Context,
		class *models.Class,
	) error
	GetAll(
		ctx context.Context,
		id string,
	) ([]*models.Class, error)
	GetClass(
		ctx context.Context,
		name string,
	) (string, error)
	GetTeacherIdbyId(
		ctx context.Context,
		id string,
	) (string,error)
Get(
		ctx context.Context,
		id string,
	) (string,error)
}

type ExamUsecase interface {
	Create(
		ctx context.Context,
		id string,
	) (
		string,
		error,
	)
	GetExist(
		ctx context.Context,
		userId string,
		id string,
	) (string,bool, error)
	MakeAttempted(
		ctx context.Context,
		id string,
		examId string,
	) error
	UpdateScore(
		ctx context.Context,
		at models.Attempted,
	) error
	GetScore(
		ctx context.Context,
		id string,
	) (int ,error)
	GetStudents(
		ctx context.Context,
		id string,
	)([]*models.Students,error)
}

type TestUsecase interface {
	CreateTest(
		ctx context.Context,
		test []*models.Tests,
	) error
	GetRandom(
		ctx context.Context,
		id string,
	) (*models.Testss, error)
	SaveAnswer(
		ctx context.Context,
		id string,
		answer models.Answer,
	) error
}
