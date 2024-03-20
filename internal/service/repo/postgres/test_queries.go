package postgres


const (
	CreateTest=`
		insert into 'test'(question,user_id) Values($1,$2) returning id
	`
	GetTest=`
	
	`
)