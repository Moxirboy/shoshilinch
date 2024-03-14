package postgres

import (
	"database/sql"
	"log"
	"errors"
	"context"
	"shoshilinch/internal/service/repo"
	"shoshilinch/internal/models"
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
	user models.User,
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
