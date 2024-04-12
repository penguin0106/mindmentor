-- +migrate Up
CREATE TABLE IF NOT EXISTS meditations (
                                           id SERIAL PRIMARY KEY,
                                           name VARCHAR(100) NOT NULL,
                                           description TEXT NOT NULL,
                                           audio_url VARCHAR(255),
                                           video_url VARCHAR(255)
);

-- +migrate Down
DROP TABLE IF EXISTS meditations;