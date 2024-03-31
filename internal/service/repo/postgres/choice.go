package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type choiceRepository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewChoiceRepository(
	db *sqlx.DB,
	log log.Logger,
) repo.ChoiceRepository {
	return &choiceRepository{
		db:  db,
		log: log,
	}
}

func (r *choiceRepository) CreateChoice(
	ctx context.Context,
	test []*models.Tests,
) error {
	r.log.Info("shuyoda")
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		r.log.Error("repo.choice.Create error while transaction begin:", err)
		return err
	}
	for i, Question := range test {
		for j, choice := range Question.Choices {
			_,err := tx.ExecContext(
				ctx,
				CreateChoice,
				Question.Test.Id,
				choice.Text,
				choice.IsTrue,
			)
			if err != nil {
				r.log.Info(err)
				_ = tx.Rollback()
				return err
			}
			err=tx.QueryRowContext(
				ctx,
				GetChoice,
				choice.Text,
			).
			Scan(
				&test[i].Choices[j].Id,
			)
			if err != nil {
				r.log.Info(err)
				_ = tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		r.log.Error("repo.choice.create error while tx commit:", err.Error())

		return err

	}
	return nil
}
