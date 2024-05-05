package repositories

import (
	"database/sql"
	"errors"
	"time"
	"trainings_service/models"
)

// CommentRepository представляет репозиторий для работы с комментариями и рейтингом тренировок
type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

// AddComment добавляет новый комментарий к тренировке в базу данных или другой источник данных
func (r *CommentRepository) AddComment(userID, trainingID int, text string) error {
	// Текущее время
	timestamp := time.Now()

	query := "INSERT INTO trainings_comments (user_id, training_id, text, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, userID, trainingID, text, timestamp)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return errors.New("ошибка при добавлении комментария")
	}

	return nil
}

// GetCommentsByTrainingID возвращает все комментарии для указанной тренировки из базы данных или другого источника данных
func (r *CommentRepository) GetCommentsByTrainingID(trainingID int) ([]*models.Comment, error) {
	query := "SELECT id, user_id, text, timestamp FROM trainings_comments WHERE training_id = $1"
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
func (r *CommentRepository) AddRating(rating *models.Rating) error {
	_, err := r.DB.Exec("INSERT INTO trainings_raiting (training_id, user_id, value) VALUES ($1, $2, $3)", rating.ItemID, rating.UserID, rating.Value)
	return err
}

// GetRating возвращает рейтинг тренировки по ее идентификатору
func (r *CommentRepository) GetRating(trainingID int) (float64, error) {
	// Реализация запроса к базе данных или другому источнику данных для получения рейтинга тренировки
	var averageRating float64
	err := r.DB.QueryRow("SELECT AVG(value) FROM trainings_raiting WHERE training_id = $1", trainingID).Scan(&averageRating)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}
