package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)


type examRepository struct{
	db *sql.DB
	log log.Logger
}

func NewExamRepository(
	db *sql.DB,
	log log.Logger,
) repo.ExamRepository {
	return &examRepository{
		log:log,
		db:db,
	}
}


func (r *examRepository) GetExist(
	ctx context.Context,
	id string,
) error {
	var exist bool
	err:=r.db.QueryRowContext(
		ctx,
		GetexistExam,
		id,
	).
	Scan(&exist)
	if err!=nil{
		r.log.Info("postgres.exam.Getexist error: ",err)
		return err
	}

	return nil
}


func (r *examRepository) Create(
	ctx context.Context,
	id string,
)( string, error){
	var examID string
	err:=r.db.QueryRowContext(
		ctx,
		id,
	).Scan(
		&examID,
	)
	if err!=nil{
		r.log.Info("poostgres.exam.Create error: ",err)
		return "", err
	}
	return examID,nil
}