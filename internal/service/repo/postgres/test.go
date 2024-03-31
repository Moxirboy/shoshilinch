package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type testRepository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewTestRepository(
	db *sqlx.DB,
	log log.Logger,
) repo.TestRepository {
	return &testRepository{
		db:  db,
		log: log,
	}
}

func (r *testRepository) Create(
	ctx context.Context,
	test []*models.Tests,
) error {
	r.log.Info("here")
	tx, err := r.db.BeginTx(
		context.Background(),
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		r.log.Error("repo.test.Create error while transaction begin:", err)
		return err
	}
	for i, Question := range test {

		_, err := tx.ExecContext(
			ctx,
			CreateTest,
			Question.Test.Question,
			Question.Test.UserID,
		)
		if err != nil {
			r.log.Info(err)
			_ = tx.Rollback()
			return err
		}
		err = tx.QueryRowContext(
			ctx,
			GetTestByQuestion,
			Question.Test.Question,
		).
			Scan(
				&test[i].Test.Id,
			)
		if err != nil {
			r.log.Info(err)
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		r.log.Error("repo.test.create error while tx commit:", err.Error())
		return err
	}

	return nil

}

func (r *testRepository) GetRandom(
	ctx context.Context,
	id string,
) (*models.Testss, error) {
	Tests := &models.Testss{}
	err := r.db.QueryRowContext(
		ctx,
		GetTest,
		id,
	).
		Scan(
			&Tests.Test.Id,
			&Tests.Test.Question,
		)
	if err != nil {
		r.log.Info("error: ", err)
		return nil, err
	}

	rows, err := r.db.QueryContext(
		ctx,
		GetChoiceByTestId,
		Tests.Test.Id,
	)
	if err != nil {
		r.log.Info("error: ", err)
		return nil, err
	}
	for rows.Next() {
		choice := models.Choices{}
		rows.Scan(
			&choice.Id,
			&choice.Text,
			&choice.IsTrue,
		)
		choice.QuestionId = Tests.Test.Id
		Tests.Choices = append(Tests.Choices, choice)
	}

	return Tests, err
}

func (r *testRepository) CreateChoice(
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
			_, err := tx.ExecContext(
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
			err = tx.QueryRowContext(
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

func (r *testRepository) SaveAnswer(
	ctx context.Context,
	id string,
	answer models.Answer,
) error {
	_, err := r.db.ExecContext(
		ctx,
		CreateAnswer,
		answer.QuestionId,
		answer.TrueAnswer,
		answer.Submitted,
		id,
	)
	if err != nil {
		r.log.Info("error: ", err)
		return err
	}
	return nil
}
