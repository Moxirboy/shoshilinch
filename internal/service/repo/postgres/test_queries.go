package postgres

const (
	CreateTest = `
		insert into 'test'(question,user_id) Values(?,?) 
	`
	GetTestByQuestion = `
	select id from 'test' where question=?
	`
	GetTest = `
	SELECT id, question
FROM 'test'

ORDER BY RANDOM();

	`

	CreateAnswer = `
	insert into 'answers'(question_id,answer,submitted,user_id) values(?,?,?,?)
`
	//WHERE user_id=?
)
