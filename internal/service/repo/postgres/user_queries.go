package postgres
const (
	// GetExist=`
	// SELECT EXIST(
	// 	SELECT * FROM user WHERE phone_number=$1
	// ) FROM user
 	// `
	GetExist=`
	SELECT EXISTS (
		SELECT * FROM user WHERE phone_number = ?
	);
	
	`
	// Create=`
	// INSERT INTO user 
	// ('phone_number','first_name','last_name','password','role')
	// VALUES($1,$2,$3,$4,$5) returning id
	// `
	Create=`
	-- Insert the data into the user table
INSERT INTO user(phone_number, first_name, last_name, password, role)
    VALUES (?, ?, ?, ?, ?);
`
GetId=`
SELECT last_insert_rowid();
`
	// Get=`
	// SELECT id,role password FROM 'user' where phone_number=$1
	// `
	Get=`
	SELECT id,role ,password FROM 'user' where phone_number=?
	`
	CleanUp=`
	Delete from 'user' where phone_number=998946320145
	`
	GetTeachers=`
	SELECT * FROM 'user' where role='teacher'`
	GetUserInfo=`
	select *from 'user' where id=?;
	`
)