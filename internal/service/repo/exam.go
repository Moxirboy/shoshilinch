package repo

import (
	"context"
	"shoshilinch/internal/models"
)


type ExamRepository interface{
	Create(
		ctx context.Context,
		id string,
	)( string, error)
	GetExist(
		ctx context.Context,
		id string,
	) (bool,error)
	MakeAttempted(
		ctx context.Context,
		id string,
		examId string,
	)error
	GetAttempted(
		ctx context.Context,
		id string,
		examId string,
	)(bool,error)
	UpdateScore(
		ctx context.Context,
		at models.Attempted,
	) error
	GetScore(
		ctx context.Context,
		id string,
	) (int, error)
	GetStudents(
		ctx context.Context,
		id string,
	)([]*models.Students,error)
	GetId(
		ctx context.Context,
		id string,
	)(ids string,err error)
}