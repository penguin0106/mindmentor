package repositories

import (
	"database/sql"
	"errors"
	"time"
	"trainings_service/models"
)

// CommentRepository представляет репозиторий для работы с комментариями и рейтингом книг
type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

// AddComment добавляет новый комментарий к книге в базу данных
func (r *CommentRepository) AddComment(userID, bookID int, text string) error {
	timestamp := time.Now()

	query := "INSERT INTO book_comments (user_id, book_id, text, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, userID, bookID, text, timestamp)
	if err != nil {
		return errors.New("ошибка при добавлении комментария к книге")
	}

	return nil
}

// GetCommentsByBookID возвращает все комментарии для указанной книги из базы данных
func (r *CommentRepository) GetCommentsByBookID(bookID int) ([]*models.Comment, error) {
	query := "SELECT id, user_id, text, timestamp FROM book_comments WHERE book_id = $1"
	rows, err := r.DB.Query(query, bookID)
	if err != nil {
		return nil, errors.New("ошибка при выполнении запроса к базе данных")
	}
	defer rows.Close()

	comments := []*models.Comment{}

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.Text, &comment.Timestamp); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// AddRating добавляет оценку для указанной книги
func (r *CommentRepository) AddRating(rating *models.Rating) error {
	_, err := r.DB.Exec("INSERT INTO book_rating (book_id, user_id, value) VALUES ($1, $2, $3)", rating.ItemID, rating.UserID, rating.Value)
	return err
}

// GetRating возвращает рейтинг книги по ее идентификатору
func (r *CommentRepository) GetRating(bookID int) (float64, error) {
	var averageRating float64
	err := r.DB.QueryRow("SELECT AVG(value) FROM book_rating WHERE book_id = $1", bookID).Scan(&averageRating)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}
