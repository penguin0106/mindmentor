package repositories

import (
	"database/sql"
	"errors"
	"mindmentor/shared/models"
)

// MusicRepository представляет собой репозиторий для работы с музыкой для медитации
type MusicRepository struct {
	DB *sql.DB
}

// GetAllMusic возвращает все аудиофайлы для медитации
func (r *MusicRepository) GetAllMusic() ([]*models.Music, error) {
	rows, err := r.DB.Query("SELECT id, name, duration, url FROM music")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var music []*models.Music
	for rows.Next() {
		var m models.Music
		if err := rows.Scan(&m.ID, &m.Name, &m.Duration, &m.URL); err != nil {
			return nil, err
		}
		music = append(music, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return music, nil
}

// GetMusicByID возвращает аудиофайл для медитации по его идентификатору
func (r *MusicRepository) GetMusicByID(musicID int) (*models.Music, error) {
	row := r.DB.QueryRow("SELECT id, name, duration, url FROM music WHERE id = $1", musicID)

	var m models.Music
	err := row.Scan(&m.ID, &m.Name, &m.Duration, &m.URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Аудиофайл с таким идентификатором не найден
		}
		return nil, err
	}

	return &m, nil
}

// AddMusic добавляет новый аудиофайл для медитации
func (r *MusicRepository) AddMusic(music *models.Music) error {
	_, err := r.DB.Exec("INSERT INTO music (name, duration, url) VALUES ($1, $2, $3)", music.Name, music.Duration, music.URL)
	return err
}
