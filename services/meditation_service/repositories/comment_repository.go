package repositories

import (
	"database/sql"
	"mindmentor/services/meditation_service/models"
)

// CommentRepository представляет собой репозиторий для работы с комментариями к курсам медитации
type CommentRepository struct {
	DB *sql.DB
}

func NewCommentrepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

// AddComment добавляет новый комментарий курса медитации
func (r *CommentRepository) AddComment(comment *models.Comment) error {
	// Реализация добавления комментария в базу данных или другой источник данных
	query := "INSERT INTO course_comments (user_id, item_id, text, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, comment.UserID, comment.ItemID, comment.Text, comment.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

// GetCommentsByCourseID возвращает все комментарии для указанного курса медитации
func (r *CommentRepository) GetCommentsByCourseID(courseID int) ([]*models.Comment, error) {
	// Реализация запроса к базе данных или другому источнику данных для получения всех комментариев курса
	query := "SELECT user_id, item_id, text, timestamp FROM course_comments WHERE item_id = $1"
	rows, err := r.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []*models.Comment{}
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.UserID, &comment.ItemID, &comment.Text, &comment.Timestamp)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

// Other methods...
