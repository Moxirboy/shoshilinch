package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
	"time"

	"github.com/jmoiron/sqlx"
)

type examRepository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewExamRepository(
	db *sqlx.DB,
	log log.Logger,
) repo.ExamRepository {
	return &examRepository{
		log: log,
		db:  db,
	}
}

func (r *examRepository) GetExist(
	ctx context.Context,
	id string,
) (bool, error) {
	var exist bool
	err := r.db.QueryRowContext(
		ctx,
		GetexistExam,
		id,
	).
		Scan(&exist)
	if err != nil {
		r.log.Info("postgres.exam.Getexist error: ", err)
		return false, err
	}

	return exist, nil
}

func (r *examRepository) GetAttempted(
	ctx context.Context,
	id string,
	examId string,
) (bool, error) {
	var attempted bool
	err := r.db.QueryRowContext(
		ctx,
		GetAttempted,
		id,
		examId,
	).
		Scan(&attempted)
	if err != nil {
		r.log.Info(err)
		return false, err
	}
	return attempted, nil
}

func (r *examRepository) Create(
	ctx context.Context,
	id string,
) (string, error) {
	_, err := r.db.ExecContext(
		ctx,
		CreateExam,
		id,
		time.Now().Format(time.RFC3339),
	)

	if err != nil {
		r.log.Info("poostgres.exam.Create error: ", err)
		return "", err
	}
	return "1", nil
}

func (r *examRepository) MakeAttempted(
	ctx context.Context,
	id string,
	examId string,
) error {
	_, err := r.db.ExecContext(
		ctx,
		attempt,
		id,
		examId,
	)
	if err != nil {
		r.log.Info("poostgres.exam.Create error: ", err)
		return err
	}
	return nil
}

func (r *examRepository) UpdateScore(
	ctx context.Context,
	at models.Attempted,
) error {
	tx, err := r.db.BeginTx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelSerializable,
		},
	)
	if err != nil {
		r.log.Error(err)
		return err
	}
	_, err = tx.ExecContext(
		ctx,
		CreateScore,
		at.Score,
		at.TeacherId,
		at.UserId,
		at.ExamId,
	)
	if err != nil {
		r.log.Info(err)
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		r.log.Info(err)
		return err
	}
	return err
}

func (r *examRepository) GetScore(
	ctx context.Context,
	id string,
) (int, error) {
	var score int
	err := r.db.QueryRowContext(
		ctx,
		GetScore,
		id,
	).
		Scan(&score)
	if err != nil {
		r.log.Info(err)
		return 0, err
	}
	return score, nil
}

func (r *examRepository) GetStudents(
	ctx context.Context,
	id string,
) ([]*models.Students, error) {
	rows, err := r.db.QueryContext(
		ctx,
		GetStudents,
		id,
	)
	if err != nil {
		r.log.Info(err)
		return nil, err
	}
	students := []*models.Students{}
	for rows.Next() {
		student := &models.Students{}
		err = rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.PhoneNumber,
			&student.Score,
		)
		if err != nil {
			r.log.Info(err)
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *examRepository) GetId(
	ctx context.Context,
	id string,
) (ids string, err error) {
	err = r.db.QueryRowContext(
		ctx,
		GetExamId,
		id,
	).
		Scan(&ids)
	if err != nil {
		r.log.Info(err)
		return "", err
	}
	return ids, nil
}
