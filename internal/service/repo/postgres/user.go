package postgres

import (
	"context"
	"database/sql"
	"errors"
	"shoshilinch/internal/models"
	"shoshilinch/internal/service/repo"
	"shoshilinch/pkg/log"

	"github.com/jmoiron/sqlx"
)


type userRepository struct{
	db *sqlx.DB
	log log.Logger
}

func NewUserRepository(
	db *sqlx.DB,
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
		GetExist,
		phoneNumber,
	).Scan(
		&exist,
	)
	r.log.Info("postgres.user.GetExist exist: ",err)
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
	r.GetTables(ctx)
	r.log.Info("postgres.user.Create user: ",user)
	_,err:=r.db.ExecContext(
		ctx,
		Create,
		user.PhoneNumber,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Role,
	)
	if err!=nil{
		r.log.Info("postgres.user.Create error: ",err)
		if err.Error()=="UNIQUE constraint failed: users.phone_number"{return "",nil}
		return "",err
	}
	err=r.db.QueryRowContext(
		ctx,
		GetId,
	).Scan(
		&id,
	)
	if err!=nil{
		r.log.Info("postgres.user.GEtId error: ",err)
		return "",err
	}
	
	return id,nil
}
func (r *userRepository) GetTables(
	ctx context.Context,
){
	rows, err := r.db.Query("SELECT name FROM sqlite_master WHERE type='table';")
    if err != nil {
        panic(err)
    }

    // Iterate over the rows and print table names
    var tableName string
    for rows.Next() {
        err := rows.Scan(&tableName)
        if err != nil {
            panic(err)
        }
        r.log.Info(tableName)
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
	defer rows.Close()
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


func (r *userRepository) CleanUp(
	ctx context.Context,
) error{
	
	_,err:=r.db.ExecContext(
		ctx,
		CleanUp,
	)
	
	if err!=nil{
		r.log.Info("postgres.user.CleanUP error: ",err)
		return err
	}
	return nil
}

func (r *userRepository) GetTeachers(
	ctx context.Context,
) ([]*models.User,error){
	var teachers []*models.User
	rows,err:=r.db.QueryContext(
		ctx,
		GetTeachers,
	)
	
	if err!=nil{
		r.log.Info("postgres.user.GetTeachers error: ",err)
		return nil,err
	}
	defer rows.Close()
	for rows.Next(){
		var teacher =&models.User{}
		err:=rows.Scan(
			&teacher.ID,
			&teacher.PhoneNumber,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Password,
			&teacher.Role,
		)
		if err!=nil{
			r.log.Info("postgres.user.GetTeachers error: ",err)
			return nil,err
		}
		teachers=append(teachers,teacher)
	}
	return teachers,nil
}


func (r *userRepository) GetUserInfo(
	ctx context.Context,
	id string,
)(*models.User,error){
	user:=&models.User{}
	err:=r.db.QueryRowContext(
		ctx,
		GetUserInfo,
		id,
	).
	Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Role,
	)
	if err!=nil{
		return nil,err
	}
	return user,nil
}