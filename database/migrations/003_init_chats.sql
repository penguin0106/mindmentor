-- +migrate Up
CREATE TABLE IF NOT EXISTS chats (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(100) NOT NULL,
                                     description TEXT NOT NULL,
                                     creator_id INT REFERENCES users(id)
);

-- +migrate Down
DROP TABLE IF EXISTS chats;