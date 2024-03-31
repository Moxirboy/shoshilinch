package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)

type classRepository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewClassRepository(
	db *sqlx.DB,
	log log.Logger,
) repo.ClassRepository {
	return &classRepository{
		db:  db,
		log: log,
	}
}
func (r *classRepository) Create(
	ctx context.Context,
	class *models.Class,
) error {
	_, err := r.db.ExecContext(
		ctx,
		CreateClass,
		class.Name,
		class.Password,
		class.UserId,
	)
	// ).
	// Scan(
	// 	&class.ID,
	// )
	if err != nil {
		r.log.Info("postgres.class.create error: ", err)
		return err
	}
	return nil
}

func (r *classRepository) GetExist(
	ctx context.Context,
	name string,
) (bool, error) {
	var exist bool
	err := r.db.QueryRowContext(
		ctx,
		GetExistClass,
		name,
	).
		Scan(
			&exist,
		)
	if err != nil {
		r.log.Info("postgres.class.getExist error: ", err)
		return false, nil
	}
	return exist, nil
}
func (r *classRepository) GetPassword(
	ctx context.Context,
	name string,
) (string, error) {
	var pass string
	err := r.db.QueryRowContext(
		ctx,
		GetPassword,
		name,
	).
		Scan(&pass)
	if err != nil {
		r.log.Info("postgres.class.GetPassword error: ", err)
		return "", err
	}
	return pass, nil
}

func (r *classRepository) GetAll(
	ctx context.Context,
	id string,
) ([]*models.Class, error) {
	rows, err := r.db.QueryContext(
		ctx,
		GetAll,
		id,
	)
	if err != nil {
		r.log.Info("postgres.class.GetAll error: ", err)
		return nil, err
	}
	classes := []*models.Class{}
	for rows.Next() {
		class := &models.Class{}
		err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.Password,
			&class.UserId,
		)
		if err != nil {
			r.log.Error("postgres.class.GetAll error: ", err)
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}
func (r *classRepository) GetId(
	ctx context.Context,
	name string,
) (string, error) {
	var id string
	err := r.db.QueryRowContext(
		ctx,
		GetClass,
		name,
	).
		Scan(
			&id,
		)
	if err != nil {
		r.log.Info("postgres.class.Get error: ", err)
		return "", err
	}
	return id, nil
}

func (r *classRepository) Get(
	ctx context.Context,
	name string,
) (string, error) {
	var id string
	err := r.db.QueryRowContext(
		ctx,
		GetClassId,
		name,
	).
		Scan(
			&id,
		)
	if err != nil {
		r.log.Info("postgres.class.Get error: ", err)
		return "", err
	}
	return id, nil
}

func (r *classRepository) BindClassStudent(
	ctx context.Context,
	classId string,
	studentId string,
) error {
	_, err := r.db.ExecContext(
		ctx,
		BindClassStudent,
		classId,
		studentId,
	)
	if err != nil {

		r.log.Info("postgres.class.BindClassStudent error: ", err)
		return err
	}
	return nil
}

func (r *classRepository) GetTeacherIdbyName(
	ctx context.Context,
	className string,
) (string,error){
	var id string

	err:=r.db.QueryRowContext(
		ctx,
		GetTeacherId,
		className,
	).
	Scan(&id)
	if err != nil {

		r.log.Info("postgres.class.getTeacherId error: ", err)
		return "",err
	}
	return id,nil
}
func (r *classRepository) GetTeacherIdbyId(
	ctx context.Context,
	classid string,
) (string,error){
	var id string

	err:=r.db.QueryRowContext(
		ctx,
		GetTeacherById,
		classid,
	).
	Scan(&id)
	if err != nil {

		r.log.Info("postgres.class.getTeacherId error: ", err)
		return "",err
	}
	return id,nil
}