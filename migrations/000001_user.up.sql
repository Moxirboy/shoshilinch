CREATE TABLE 'user' {
    "id" SERIAL PRIMARY KEY,
    "phone_number" text NOT NULL,
    "first_name"  VARCHAR(45) NOT NULL,
    "last_name"   VARCHAR(45) NOT NULL,
    "password"   VARCHAR(45) UNIQUE NOT NULL,
    "role" role NOT NULL,
}

CREATE TYPE role AS enum ('teacher','student','admin');