CREATE TABLE "choice"(
    'id' SERIAL PRIMARY KEY,
    'question_id' VARCHAR(40),
    'text' TEXT,
    'is_true' boolean
);