package postgres

const (
	CreateAuth=`
	INSERT INTO 'auth'
	('id','role','token','datetime') 
	VALUES($1,$2,$3,$4)
	`
	GetExistToken=`
	SELECT EXIST(
		SELECT *
		FROM 'auth' WHERE token=$1 AND datetime::timestamp>$2::timestamp
	) FROM 'auth'
	`
	authGet=`
	SELECT 
	id,role,token,datetime 
	FROM 'auth' WHERE 
	token = $1 AND datetime::timestamp > $2::timestamp
	`
	cleanUp=`
	Delete FROM 'auth' WHERE datetime < CURRENT_TIMESTAMP
	`
	authDelete=`
		DELETE FROM 'auth' where token=$1
	`
)