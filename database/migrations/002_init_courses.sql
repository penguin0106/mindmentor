-- +migrate Up
CREATE TABLE IF NOT EXISTS courses (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(100) NOT NULL,
                                       description TEXT NOT NULL,
                                       author_id INT REFERENCES users(id)
);

-- +migrate Down
DROP TABLE IF EXISTS courses;