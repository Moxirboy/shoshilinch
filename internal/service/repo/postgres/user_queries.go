package postgres
const (
	GetExist=`
	SELECT EXIST(
		SELECT * FROM user WHERE phone_number=$1
	) FROM user
 	`
	Create=`
	INSERT INTO user 
	('phone_number','first_name','last_name','password','role')
	VALUES($1,$2,$3,$4,$5) returning id
	`
)