package repositories

import (
	"database/sql"
	"errors"
	"mindmentor/services/emotions_service/models"
	"time"
)

// EmotionRepository представляет собой репозиторий для работы с записями эмоций
type EmotionRepository struct {
	DB *sql.DB
}

// CreateEmotion создает новую запись эмоции в базе данных
func (r *EmotionRepository) CreateEmotion(emotion *models.Emotion) error {
	query := "INSERT INTO emotions (topic, body, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	now := time.Now()
	_, err := r.DB.Exec(query, emotion.Topic, emotion.Body, emotion.UserID, now, now)
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
func (r *EmotionRepository) GetEmotionsByUserID(userID int) ([]*models.Emotion, error) {
	// Запрос к базе данных для выборки эмоций пользователя по его идентификатору
	query := "SELECT id, topic, body, user_id, created_at FROM emotions WHERE user_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Слайс для хранения эмоций пользователя
	emotions := []*models.Emotion{}

	// Итерация по результатам запроса
	for rows.Next() {
		var emotion models.Emotion
		err := rows.Scan(&emotion.ID, &emotion.Topic, &emotion.Body, &emotion.UserID, &emotion.CreatedAt)
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
