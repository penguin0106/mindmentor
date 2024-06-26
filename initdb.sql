CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(255) NOT NULL UNIQUE,
                                     email VARCHAR(255) NOT NULL UNIQUE,
                                     password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS meditation_videos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    video_content BYTEA NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS meditation_music (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     duration INT NOT NULL,
                                     music_content BYTEA NOT NULL,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS video_comments (
                                               id SERIAL PRIMARY KEY,
                                               user_id INT NOT NULL,
                                               item_id INT NOT NULL,
                                               text TEXT,
                                               timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                               CONSTRAINT fk_user_video_comments FOREIGN KEY (user_id) REFERENCES users(id),
                                               CONSTRAINT fk_video_comments FOREIGN KEY (item_id) REFERENCES meditation_videos(id)
);

CREATE TABLE IF NOT EXISTS video_ratings (
                                              id SERIAL PRIMARY KEY,
                                              video_id INT NOT NULL,
                                              user_id INT NOT NULL,
                                              rating FLOAT,
                                              created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                              CONSTRAINT fk_user_video_rating FOREIGN KEY (user_id) REFERENCES users(id),
                                              CONSTRAINT fk_video_video_rating FOREIGN KEY (video_id) REFERENCES meditation_videos(id)
);

CREATE TABLE IF NOT EXISTS video_favorites (
                                                id SERIAL PRIMARY KEY,
                                                user_id INT NOT NULL,
                                                item_id INT NOT NULL,
                                                CONSTRAINT fk_user_course_favorite FOREIGN KEY (user_id) REFERENCES users(id),
                                                CONSTRAINT fk_video_favorite FOREIGN KEY (item_id) REFERENCES meditation_videos(id)
);

CREATE TABLE IF NOT EXISTS emotions (
                                        id SERIAL PRIMARY KEY,
                                        topic VARCHAR(255) NOT NULL,
                                        body TEXT

);

CREATE TABLE IF NOT EXISTS books (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     description TEXT,
                                     content BYTEA NOT NULL,
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS book_comments (
                                             id SERIAL PRIMARY KEY,
                                             user_id INT NOT NULL,
                                             book_id INT NOT NULL,
                                             text TEXT NOT NULL,
                                             timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                             FOREIGN KEY (user_id) REFERENCES users(id),
                                             FOREIGN KEY (book_id) REFERENCES books(id)
);

CREATE TABLE IF NOT EXISTS book_rating (
                                           id SERIAL PRIMARY KEY,
                                           user_id INT NOT NULL,
                                           book_id INT NOT NULL,
                                           value FLOAT NOT NULL,
                                           FOREIGN KEY (user_id) REFERENCES users(id),
                                           FOREIGN KEY (book_id) REFERENCES books(id)
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

CREATE TABLE IF NOT EXISTS messages (
                                        id SERIAL PRIMARY KEY,
                                        discussion_id INT NOT NULL,
                                        user_id INT NOT NULL,
                                        text TEXT NOT NULL,
                                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                        FOREIGN KEY (discussion_id) REFERENCES discussions(id) ON DELETE CASCADE,
                                        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);