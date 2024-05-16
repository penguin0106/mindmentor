package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"meditation_service/models"
)

// MusicRepository представляет собой репозиторий для работы с музыкой для медитации
type MusicRepository struct {
	DB *sql.DB
}

func NewMusicRepository(db *sql.DB) *MusicRepository {
	return &MusicRepository{DB: db}
}

// AddMusic добавляет новый аудиофайл для медитации
func (r *MusicRepository) AddMusic(music *models.Music) error {
	_, err := r.DB.Exec("INSERT INTO meditation_music (name, duration, music_content) VALUES ($1, $2, $3)", music.Name, music.Duration, music.MusicFile)
	return err
}

// GetMusicByID возвращает аудиофайл для медитации по его идентификатору
func (r *MusicRepository) GetMusicByID(musicID int) (*models.Music, error) {
	row := r.DB.QueryRow("SELECT id, name, duration, music_content, created_at FROM meditation_music WHERE id = $1", musicID)

	var music models.Music
	err := row.Scan(&music.ID, &music.Name, &music.Duration, &music.MusicFile, &music.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &music, nil
}

// GetAllMusic возвращает все аудиофайлы для медитации
func (r *MusicRepository) GetAllMusic() ([]*models.Music, error) {
	rows, err := r.DB.Query("SELECT id, name, duration, created_at FROM meditation_music")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var music []*models.Music
	for rows.Next() {
		var m models.Music
		if err := rows.Scan(&m.ID, &m.Name, &m.Duration, &m.CreatedAt); err != nil {
			return nil, err
		}
		music = append(music, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return music, nil
}
