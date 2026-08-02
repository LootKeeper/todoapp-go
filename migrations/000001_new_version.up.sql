CREATE SCHEMA todoapp;

CREATE TABLE todoapp.users (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  name VARCHAR(100) NOT NULL CHECK(char_length(name) BETWEEN 3 AND 100),
  email VARCHAR(150) NOT NULL,
  phone_number VARCHAR(15) CHECK(
    phone_number ~ '^\+[0-9]+$'
    AND
    char_length(phone_number) BETWEEN 10 AND 15
  )
);

CREATE TABLE todoapp.tasks (
  id SERIAL PRIMARY KEY,
  version BIGINT NOT NULL DEFAULT 1,
  title VARCHAR(200) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 200),
  description VARCHAR(2000),
  completed BOOLEAN NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  completed_at TIMESTAMPTZ,

  CHECK (
    (completed=FALSE AND completed_at IS NULL)
    OR 
    (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
  ),

  author_user_id INTEGER NOT NULL REFERENCES todoapp.users(id)
)