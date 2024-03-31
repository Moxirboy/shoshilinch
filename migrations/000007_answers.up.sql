CREATE TABLE IF NOT EXISTs 'answers'(
    "id" SERIAL PRIMARY KEY,
    "question_id" VARCHAR(60),
    "answer" Text,
    "submitted" Text,
    "user_id" INT NOT NULL,
)