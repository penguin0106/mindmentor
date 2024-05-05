package repositories

import (
	"context"
	"database/sql"
	"social_service/models"
	"time"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

// CreateMessage создает новое сообщение в указанном чате.
func (r *MessageRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	query := `INSERT INTO messages (discussion_id, user_id, text, created_at)
              VALUES ($1, $2, $3, $4)`

	_, err := r.DB.ExecContext(ctx, query, message.DiscussionID, message.UserID, message.Text, message.CreationTime)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMessage обновляет существующее сообщение в указанном чате.
func (r *MessageRepository) UpdateMessage(ctx context.Context, message *models.Message) error {
	query := `UPDATE messages
              SET text = $1, updated_at = $2
              WHERE id = $3`

	_, err := r.DB.ExecContext(ctx, query, message.Text, time.Now(), message.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMessage удаляет сообщение из указанного чата.
func (r *MessageRepository) DeleteMessage(ctx context.Context, messageID int) error {
	query := `DELETE FROM messages WHERE id = $1`

	_, err := r.DB.ExecContext(ctx, query, messageID)
	if err != nil {
		return err
	}

	return nil
}

// GetMessageByID возвращает сообщение по его идентификатору.
func (r *MessageRepository) GetMessageByID(ctx context.Context, messageID int) (*models.Message, error) {
	query := `SELECT id, discussion_id, user_id, text, created_at, updated_at
              FROM messages
              WHERE id = $1`

	row := r.DB.QueryRowContext(ctx, query, messageID)

	var message models.Message
	err := row.Scan(&message.ID, &message.DiscussionID, &message.UserID, &message.Text, &message.CreationTime, &message.LastEditTime)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// GetMessagesByDiscussion возвращает все сообщения в указанном чате.
func (r *MessageRepository) GetMessagesByDiscussion(ctx context.Context, discussionID int) ([]*models.Message, error) {
	query := `SELECT id, user_id, text, created_at, updated_at
              FROM messages
              WHERE discussion_id = $1`

	rows, err := r.DB.QueryContext(ctx, query, discussionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.UserID, &message.Text, &message.CreationTime, &message.LastEditTime); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
