CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS music (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    duration INT NOT NULL,
    url VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS course_comments (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  course_id INT NOT NULL,
  text TEXT,
  timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_course_comments FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_course_comments FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE IF NOT EXISTS course_ratings (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    user_id INT NOT NULL,
    value FLOAT,
    CONSTRAINT fk_user_course_rating FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_course_course_rating FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE IF NOT EXISTS course_favorites (
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   course_id INT NOT NULL,
   CONSTRAINT fk_user_course_favorite FOREIGN KEY (user_id) REFERENCES users(id),
   CONSTRAINT fk_course_favorite FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE IF NOT EXISTS emotions (
    id SERIAL PRIMARY KEY,
    topic VARCHAR(255) NOT NULL,
    body TEXT,
    user_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_emotions FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS trainings (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  rating FLOAT,
  favorite BOOLEAN
);

CREATE TABLE IF NOT EXISTS trainings_comments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    training_id INT NOT NULL,
    text TEXT,
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_trainings_comments FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_trainings_comments FOREIGN KEY (training_id) REFERENCES trainings(id)
);

CREATE TABLE IF NOT EXISTS trainings_rating (
    id SERIAL PRIMARY KEY,
    training_id INT NOT NULL,
    user_id INT NOT NULL,
    value FLOAT,
    CONSTRAINT fk_user_rating FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_training_rating FOREIGN KEY (training_id) REFERENCES trainings(id)
);

CREATE TABLE IF NOT EXISTS trainings_favorites (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    training_id INT NOT NULL,
    CONSTRAINT fk_user_training_favorite FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_training_favorite FOREIGN KEY (training_id) REFERENCES trainings(id)
);

CREATE TABLE IF NOT EXISTS discussions (
    id SERIAL PRIMARY KEY,
    topic VARCHAR(255) NOT NULL,
    owner_id INT NOT NULL,
    CONSTRAINT fk_user_discussions FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS user_discussions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    discussion_id INT NOT NULL,
    CONSTRAINT fk_user_discussion FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_discussion_user FOREIGN KEY (discussion_id) REFERENCES discussions(id)
);