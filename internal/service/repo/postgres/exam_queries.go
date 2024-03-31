package postgres



const(
	// GetexistExam=`
	// select Exist (
	// 	select *from exam where class_id=$1 AND created_at+15<CURRENT_TIMESTAMP
	// ) from exam
	// `
	GetexistExam=`
	SELECT EXISTS(SELECT 1 FROM exam WHERE class_id = ? Order by id desc limit 1 )
	`
	GetExamId=`
	SELECT id from exam where class_id=? order by id desc limit 1
	`
	GetAttempted=`
	Select EXISTS(select 1 from attempt where user_id=? and complete=1  and exam_id=? )
	`
	// CreateExam=`
	// insert into exam (class_id,created_at) values($1,$2) returning id
	// `
	CreateExam=`
	insert into exam (class_id,created_at) values(?,?)
	`
	attempt=`
	insert into attempt (user_id,exam_id) values(?,?)
	`
	CreateScore=`
	UPDATE attempt
SET score = ?,teacher_id=?, complete = 1
WHERE user_id = ? and exam_id=?;

	`
	GetScore=`
	select score from attempt where user_id=? ORDER BY id DESC LIMIT 1;
	`
	GetStudents=`
	SELECT u.id, u.first_name, u.last_name, u.phone_number, a.score
FROM "user" u
JOIN attempt a ON u.id = a.user_id
WHERE a.teacher_id = ?

	`
)