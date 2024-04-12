-- +migrate Up
CREATE TABLE IF NOT EXISTS emotions (
                                        id SERIAL PRIMARY KEY,
                                        user_id INT REFERENCES users(id),
                                        date DATE NOT NULL,
                                        emotion VARCHAR(50) NOT NULL,
                                        mood VARCHAR(50) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS emotions;