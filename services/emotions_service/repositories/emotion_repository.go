package repositories

import (
	"database/sql"
	"emotions_service/models"
	"errors"
	"time"
)

// EmotionRepository представляет собой репозиторий для работы с записями эмоций
type EmotionRepository struct {
	DB *sql.DB
}

func NewEmotionRepository(db *sql.DB) *EmotionRepository {
	return &EmotionRepository{DB: db}
}

// CreateEmotion создает новую запись эмоции в базе данных
func (r *EmotionRepository) CreateEmotion(emotion *models.Emotion) error {
	query := "INSERT INTO emotions (topic, body) VALUES ($1, $2)"

	_, err := r.DB.Exec(query, emotion.Topic, emotion.Body)
	if err != nil {
		return err
	}
	return nil
}

// UpdateEmotion обновляет существующую запись эмоции в базе данных
func (r *EmotionRepository) UpdateEmotion(emotionID int, updatedEmotion *models.Emotion) error {
	query := "UPDATE emotions SET topic = $1, body = $2, updated_at = $3 WHERE id = $4"
	now := time.Now()
	result, err := r.DB.Exec(query, updatedEmotion.Topic, updatedEmotion.Body, now, emotionID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("запись эмоции с указанным ID не найдена")
	}
	return nil
}

// DeleteEmotion удаляет запись эмоции из базы данных
func (r *EmotionRepository) DeleteEmotion(emotionID int) error {
	query := "DELETE FROM emotions WHERE id = $1"
	result, err := r.DB.Exec(query, emotionID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("запись эмоции с указанным ID не найдена")
	}
	return nil
}

// GetEmotionsByUserID возвращает эмоции пользователя по его идентификатору
func (r *EmotionRepository) GetEmotionsByUserID() ([]*models.Emotion, error) {
	// Запрос к базе данных для выборки эмоций пользователя по его идентификатору
	query := "SELECT id, topic, body FROM emotions"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Слайс для хранения эмоций пользователя
	var emotions []*models.Emotion

	// Итерация по результатам запроса
	for rows.Next() {
		var emotion models.Emotion
		err := rows.Scan(&emotion.ID, &emotion.Topic, &emotion.Body)
		if err != nil {
			return nil, err
		}
		emotions = append(emotions, &emotion)
	}

	// Проверка наличия ошибок при итерации
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return emotions, nil
}
