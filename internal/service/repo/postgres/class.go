package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
)



type classRepository struct{
	db *sql.DB
	log log.Logger
}

func NewClassRepository(
	db *sql.DB,
	log log.Logger,
) repo.ClassRepository{
	return &classRepository{
		db: db,
		log: log,
	}
}
func (r *classRepository) Create(
	ctx context.Context,
	class *models.Class,
)error{
	err:=r.db.QueryRowContext(
		ctx,
		class.Name,
		class.Password,
		class.UserId,
	).
	Scan(
		&class.ID,
	)
	if err!=nil{
		r.log.Info("postgres.class.create error: ",err)
		return err
	}
	return nil
}

func (r *classRepository) GetExist(
	ctx context.Context,
	name string,
) (bool,error){
	var exist bool
		err:=r.db.QueryRowContext(
			ctx,
			GetExistClass,
			name,
		).
		Scan(
			&exist,
		)
	if err!=nil{
		r.log.Info("postgres.class.getExist error: ",err)
		return false,nil
	}
	return exist,nil
}
func (r *classRepository) GetPassword(
	ctx context.Context,
	name string,
) (string,error){
	var pass string
	err:=r.db.QueryRowContext(
		ctx,
		GetPassword,
		name,
	).
	Scan(&pass)
	if err!=nil{
		r.log.Info("postgres.class.GetPassword error: ",err)
		return "",err
	}
	return pass,nil
}

func (r *classRepository) GetAll(
	ctx context.Context,
	id string,
) ([]*models.Class,error){
	rows,err:=r.db.QueryContext(
		ctx,
		GetAll,
		id,
	)
	if err!=nil{
		r.log.Info("postgres.class.GetAll error: ",err)
		return nil,err
	}
	classes:=[]*models.Class{}
	for rows.Next(){
		class:=&models.Class{}
		err:=rows.Scan(
			&class.ID,
			&class.Name,
			&class.Password,
			&class.UserId,
		)
		if err!=nil{
			r.log.Error("postgres.class.GetAll error: ",err)
			return nil,err
		}
		classes=append(classes, class)
	}
	return classes,nil
}
