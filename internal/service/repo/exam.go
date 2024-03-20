package repo

import "context"


type ExamRepository interface{
	Create(
		ctx context.Context,
		id string,
	)( string, error)
	GetExist(
		ctx context.Context,
		id string,
	) error
}