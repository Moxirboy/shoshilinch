package repo
import (
	"context"
	"shoshilinch/internal/models"
)

type ClassRepository interface{
	GetExist(
		ctx context.Context,
		name string,
	) (bool,error)
	GetPassword(
		ctx context.Context,
		name string,
	) (string,error)
	Create(
		ctx context.Context,
		class *models.Class,
	)error
	GetAll(
		ctx context.Context,
		id string,
	) ([]*models.Class,error)
	Get(
		ctx context.Context,
		name string,
	) (string,error)
	BindClassStudent(
		ctx context.Context,
		classId string,
		studentId string,
	) error
	GetTeacherIdbyName(
		ctx context.Context,
		className string,
	) (string,error)
	GetTeacherIdbyId(
		ctx context.Context,
		classid string,
	) (string,error)
	GetId(
		ctx context.Context,
		name string,
	) (string, error)
}