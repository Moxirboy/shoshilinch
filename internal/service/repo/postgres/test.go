package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)


type testRepository struct{
	db *sql.DB
	log log.Logger
}

func NewTestRepository(
	db *sql.DB,
	log log.Logger,
) repo.TestRepository{
	return &testRepository{
		db: db,
		log: log,
	}
}

func (r *testRepository) Create(
	ctx context.Context,
	test []*models.Tests,
)error{
	tx,err:=r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err!=nil{
		r.log.Error("repo.test.Create error while transaction begin:", err)
		return err
	}
	for i,Question:=range test{

		
		err:=tx.QueryRowContext(
			ctx,
			CreateTest,
			Question.Test.Question,
			Question.Test.UserID,
		).
		Scan(
			&test[i].Test.Id,
		)
		if err!=nil{
			_=tx.Rollback()
			return err 
			}
		}
	

	if err := tx.Commit(); err != nil {
		r.log.Error("repo.test.create error while tx commit:", err.Error())
		return err
	}
	
	return nil

}