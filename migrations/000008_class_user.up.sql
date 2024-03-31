CREATE TABLE IF NOT EXISTS "class_user" (
  "id" SERIAL PRIMARY KEY,
  "class_id" VARCHAR(60),
  "student_id"  VARCHAR(60)
);