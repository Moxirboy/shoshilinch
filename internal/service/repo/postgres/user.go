package postgres

import (
	"context"
	"database/sql"
	"errors"
	"shoshilinch/pkg/log"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
)


type userRepository struct{
	db *sql.DB
	log log.Logger
}

func NewUserRepository(
	db *sql.DB,
	log log.Logger,
) repo.UserRepository{
	return &userRepository{
		db:db,
		log:log,
	}
}
func (r *userRepository) GetExist(
	ctx context.Context,
	phoneNumber string,
	) (
	bool,
	error,
){
	var exist bool
	err:=r.db.QueryRowContext(
		ctx,
		phoneNumber,
	).Scan(
		&exist,
	)
	if errors.Is(err,sql.ErrNoRows){
		return false,nil
	}else if err!=nil {
		return false,nil
	}
	return exist,nil
}

func (r *userRepository) Create(
	ctx context.Context,
	user *models.User,
) (string,error){
	var id string
	err:=r.db.QueryRowContext(
		ctx,
		Create,
		user.PhoneNumber,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Role,
	).Scan(
		&id,
	)
	if err!=nil{
		return "",err
	}
	return id,nil
}
func(r *userRepository) Get(
	ctx context.Context,
	phoneNumber string,
) (string,string,string,error){
	var password string
	var id string
	var role string
	err:=r.db.QueryRowContext(
		ctx,
		Get,
		phoneNumber,
	).
	Scan(
		&id,
		&role,
		&password,
	)
	if err!=nil{
		r.log.Info("postgres.user.Get error: ",err)
		return "","","",err
	}
	return id,role,password,nil
}