package postgres

import (
	"context"
	"database/sql"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"
	"time"
)


type authRepository struct{
	db *sql.DB
	log log.Logger
}

func NewAuthRepository(
	db *sql.DB,
	log log.Logger,
) repo.AuthRepository{
	return &authRepository{
		db:db,
		log:log,
	}
}


func (r *authRepository) Create(
	ctx context.Context,
	token string,
	id string,
	role string,
	date time.Time,
)error{	
	_,err:=r.db.ExecContext(
		ctx,
		CreateAuth,
		id,
		role,
		token,
		date.Format(time.RFC3339),
	)
	if err!=nil{
		r.log.Info("postgres.authCreate error: ",err)
		return err
	}
	return nil
}

func(r *authRepository) GetExist(
	ctx context.Context,
	token string,
	date time.Time,
) (
	bool,
	error,
){
	var exist bool
	err:=r.db.QueryRowContext(
		ctx,
		token,
		date.Format(time.RFC3339),
	).
	Scan(
		&exist,
	)
	if err!=nil{
		return false,err
	}
	return exist,nil
}

func (r *authRepository) Get(
	ctx context.Context,
	token string,
	date time.Time,
) (
	*models.Token,
	error,
){
	at:=&models.Token{}
	err:=r.db.QueryRowContext(
		ctx,
		authGet,token,
		date.Format(time.RFC3339),
	).
	Scan(
		&at.Id,
		&at.Role,
		&at.Token,
		&at.Datetime,
	)
	if err!=nil{
		r.log.Info("postgres.Get error: ",err)
		return nil,err
	}
	return at,nil
}
func (r *authRepository) CleanUp(
	ctx context.Context,
	
)error{
	_,err:=r.db.ExecContext(
		ctx,
		cleanUp,
	)
	if err!=nil{
		r.log.Info("postgres.cleanup error: ",err)
		return err
	}
	return nil
}

func (r *authRepository) Delete(
	ctx context.Context,
	token string,
) error{
	_,err:=r.db.ExecContext(
		ctx,
		authDelete,
		token,
	)
	if err!=nil{
		return err
	}
	return nil
}