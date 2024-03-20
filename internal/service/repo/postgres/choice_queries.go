package postgres

const (
	CreateChoice = `
		insert into 'choice'(question_id,text,is_true) values($1,$2,$3) returning id
	`
)
