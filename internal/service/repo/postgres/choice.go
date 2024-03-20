package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type choiceRepository struct {
	db  *sql.DB
	log log.Logger
}

func NewChoiceRepository(
	db *sql.DB,
	log log.Logger,
) repo.ChoiceRepository {
	return &choiceRepository{
		db:  db,
		log: log,
	}
}

func (r *choiceRepository) Create(
	ctx context.Context,
	test []*models.Tests,
) error {
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
			err := tx.QueryRowContext(
				ctx,
				CreateChoice,
				Question.Test.Id,
				choice.Text,
				choice.IsTrue,
			).
				Scan(
					&test[i].Choices[j].Id,
				)
			if err != nil {
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
