package repositories

import (
	"database/sql"
	"errors"
	"mindmentor/services/trainings_service/models"
	"time"
)

// CommentRepository представляет репозиторий для работы с комментариями и рейтингом тренировок
type CommentRepository struct {
	DB *sql.DB
}

// AddComment добавляет новый комментарий к тренировке в базу данных или другой источник данных
func (r *CommentRepository) AddComment(userID, trainingID int, text string) error {
	// Текущее время
	timestamp := time.Now()

	query := "INSERT INTO comments (user_id, training_id, text, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, userID, trainingID, text, timestamp)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return errors.New("ошибка при добавлении комментария")
	}

	return nil
}

// GetCommentsByTrainingID возвращает все комментарии для указанной тренировки из базы данных или другого источника данных
func (r *CommentRepository) GetCommentsByTrainingID(trainingID int) ([]*models.Comment, error) {
	query := "SELECT id, user_id, text, timestamp FROM comments WHERE training_id = $1"
	rows, err := r.DB.Query(query, trainingID)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return nil, errors.New("ошибка при выполнении запроса к базе данных")
	}
	defer rows.Close()

	// Инициализируем слайс для хранения комментариев
	comments := []*models.Comment{}

	// Обходим результаты запроса
	for rows.Next() {
		var comment models.Comment
		// Сканируем строки и создаем объекты комментариев
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.Text, &comment.Timestamp); err != nil {
			// Возвращаем ошибку, если произошла ошибка при сканировании строки
			return nil, err
		}
		// Добавляем комментарий в слайс
		comments = append(comments, &comment)
	}

	// Проверяем наличие ошибок после обхода результатов
	if err := rows.Err(); err != nil {
		// Возвращаем ошибку, если произошла ошибка после обхода результатов
		return nil, err
	}

	// Возвращаем слайс комментариев и nil в качестве ошибки
	return comments, nil
}

// AddRating добавляет оценку для указанной тренировки
func (r *CommentRepository) AddRating(trainingID int, rating float64) error {
	query := "UPDATE trainings SET rating = $1 WHERE id = $2"
	result, err := r.DB.Exec(query, rating, trainingID)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return err
	}

	// Проверяем количество обновленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Возвращаем ошибку, если не удалось получить количество обновленных строк
		return err
	}
	if rowsAffected == 0 {
		// Возвращаем ошибку, если не удалось найти тренировку с указанным ID
		return errors.New("тренировка с указанным ID не найдена")
	}

	return nil
}

// GetRating возвращает рейтинг тренировки по ее идентификатору
func (r *CommentRepository) GetRating(trainingID int) (float64, error) {
	// Реализация запроса к базе данных или другому источнику данных для получения рейтинга тренировки
	query := "SELECT rating FROM trainings WHERE id = $1"
	row := r.DB.QueryRow(query, trainingID)

	var rating float64
	err := row.Scan(&rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Тренировка с указанным ID не найдена
			return 0, errors.New("тренировка с указанным ID не найдена")
		}
		return 0, err
	}

	return rating, nil
}
