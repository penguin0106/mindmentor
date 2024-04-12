-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(50) UNIQUE NOT NULL,
                                     email VARCHAR(100) UNIQUE NOT NULL,
                                     password VARCHAR(100) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;