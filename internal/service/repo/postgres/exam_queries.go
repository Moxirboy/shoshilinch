package postgres



const(
	GetexistExam=`
	select Exist (
		select *from exam where class_id=$1 AND created_at+15<CURRENT_TIMESTAMP
	) from exam
	`
	CreateExam=`
	insert into exam (class_id,created_at) values($1,$2) returning id
	`
)