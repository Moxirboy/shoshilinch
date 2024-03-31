create TABLE IF NOT EXISTS attempt (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    teacher_id INT ,
    score INT ,
    exam_id INT ,
    complete int NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
