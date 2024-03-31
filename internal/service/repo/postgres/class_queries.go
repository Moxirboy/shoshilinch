package postgres

const (
	GetPassword = `
	SELECT password from 'class' where name=$1
	`
	// GetExistClass=`
	// SELECT EXISTS(
	// 	SELECT * FROM class WHERE name=$1
	// ) FROM class
	// `
	GetExistClass = `
	 SELECT EXISTS(
		 SELECT 1 FROM class WHERE name=?
	 ) 
	  `
	// CreateClass=`
	// INSERT INTO 'class' (name,password,teacher_id) VALUES($1,$2,$3) returning id
	// `
	CreateClass = `
	INSERT INTO class (name,password,teacher_id) VALUES(?,?,?) 
	`
	GetAll = `
	Select * from 'class' where teacher_id=$1
	`
	GetClassId = `
	SELECT class_id from 'class_user' where student_id=?`
	GetClass = `
	SELECT id from 'class' where name=?`
	BindClassStudent = ` 
	INSERT INTO 'class_user' (class_id,student_id) VALUES($1,$2)
	`
	GetTeacherId=`
	select teacher_id from 'class' where name=?
	`
	GetTeacherById=`
	select teacher_id from 'class' where id=?
	`
)
