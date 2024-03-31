package postgres

const (
	CreateChoice = `
		insert into 'choice'(question_id,text,is_true) values(?,?,?) 
	`
	GetChoice=`
	select id from 'choice' where text=?
	`
	GetChoiceByTestId=`
		select id,text,is_true from choice where question_id=?
	`
)