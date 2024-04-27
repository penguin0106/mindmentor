package repositories

import (
	"database/sql"
	"mindmentor/services/emotions_service/models"
)

// EmotionRepository представляет собой репозиторий для работы с эмоциями (записями)
type EmotionRepository struct {
	DB *sql.DB
}

// CreateEmotion создает новую эмоцию (запись) в базе данных
func (r *EmotionRepository) CreateEmotion(emotion *models.Emotion) error {
	// Подготовка запроса на вставку новой эмоции
	query := "INSERT INTO emotions (topic, body, user_id) VALUES ($1, $2, $3)"
	// Выполнение запроса к базе данных
	_, err := r.DB.Exec(query, emotion.Topic, emotion.Body, emotion.UserID)
	if err != nil {
		// В случае ошибки вернуть ошибку
		return err
	}
	// Возвращаем nil, если операция выполнена успешно
	return nil
}

// GetEmotionsByUserID возвращает все эмоции (записи) для указанного пользователя
func (r *EmotionRepository) GetEmotionsByUserID(userID int) ([]*models.Emotion, error) {
	// Подготовка запроса на выборку эмоций для указанного пользователя
	query := "SELECT id, topic, body, user_id FROM emotions WHERE user_id = $1"
	// Выполнение запроса к базе данных
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		// В случае ошибки вернуть ошибку
		return nil, err
	}
	defer rows.Close()

	// Создаем слайс для хранения эмоций пользователя
	emotions := []*models.Emotion{}

	// Итерируем по результатам запроса и добавляем эмоции в слайс
	for rows.Next() {
		var emotion models.Emotion
		// Сканируем строки результата запроса в структуру эмоции
		err := rows.Scan(&emotion.ID, &emotion.Topic, &emotion.Body, &emotion.UserID)
		if err != nil {
			// В случае ошибки вернуть ошибку
			return nil, err
		}
		// Добавляем эмоцию в слайс
		emotions = append(emotions, &emotion)
	}

	// Проверяем наличие ошибок после итерации по результатам запроса
	if err := rows.Err(); err != nil {
		// В случае ошибки вернуть ошибку
		return nil, err
	}

	// Возвращаем слайс эмоций пользователя и nil в качестве ошибки
	return emotions, nil
}
