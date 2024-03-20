package postgres


const (
	GetPassword=`
	SELECT password from 'class' where name=$1
	`
	GetExistClass=`
	SELECT EXIST(
		SELECT * FROM class WHERE name=$1
	) FROM class
 	`
	CreateClass=`
	INSERT INTO 'class' (name,password,teacher_id) VALUES($1,$2,$3) returning id
	`
	GetAll=`
	Select * from 'class' where teacher_id=$1
	`
)