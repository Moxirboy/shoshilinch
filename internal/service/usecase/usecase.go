package usecase
import (
	"context"
	"shoshilinch/internal/models"

)
type AuthUseCase interface{
	New(
		ctx context.Context,
		id,role string,
	)(
		string,
		error,
	)
	Check(
		ctx context.Context,
		token string,
	)(id ,role string,err error)
}


type UserUsecase interface{
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
	)(string,string,error)
	SignUp(
		ctx context.Context,
		user *models.User,
		className string,
	)(
		string,
		error,
	)
}


type ClassUsecase interface{
	Create(
		ctx context.Context,
		class *models.Class,
	) error
	GetAll(
		ctx context.Context,
		id string,
	) ([]*models.Class,error)
}

type ExamUsecase interface{
	Create(
		ctx context.Context,
		id string,
	)( 
	string, 
	error,
)
	GetExist(
		ctx context.Context,
		id string,
	) error
}

type TestUsecase interface{
	CreateTest(
		ctx context.Context,
		test []*models.Tests,
	) error
}