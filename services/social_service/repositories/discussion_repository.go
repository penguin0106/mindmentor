package repositories

import (
	"context"
	"database/sql"
	"errors"
	"mindmentor/shared/models"
	"time"
)

type DiscussionRepository struct {
	DB *sql.DB
}

func NewDiscussionRepository(db *sql.DB) *DiscussionRepository {
	return &DiscussionRepository{DB: db}
}

// CreateDiscussion создает новое обсуждение в базе данных
func (r *DiscussionRepository) CreateDiscussion(discussion *models.Discussion) error {
	query := "INSERT INTO discussions (topic, owner_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, discussion.Topic, discussion.OwnerID)
	if err != nil {
		return err
	}
	return nil
}

// FindDiscussion ищет обсуждение по его теме в базе данных
func (r *DiscussionRepository) FindDiscussion(topic string) (*models.Discussion, error) {
	query := "SELECT id, topic, owner_id FROM discussions WHERE topic = $1"
	row := r.DB.QueryRow(query, topic)

	var discussion models.Discussion
	err := row.Scan(&discussion.ID, &discussion.Topic, &discussion.OwnerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Обсуждение с указанной темой не найдено
			return nil, nil
		}
		return nil, err
	}

	return &discussion, nil
}

// JoinDiscussion добавляет пользователя к обсуждению в базе данных
func (r *DiscussionRepository) JoinDiscussion(ctx context.Context, userID, discussionID int) error {
	query := "INSERT INTO user_discussions (user_id, discussion_id) VALUES ($1, $2)"
	_, err := r.DB.ExecContext(ctx, query, userID, discussionID)
	if err != nil {
		return err
	}

	// Получаем сообщения пользователя
	userMessages := getUsersMessages(userID)

	// Добавляем каждое сообщение пользователя к обсуждению
	for _, message := range userMessages {
		message.DiscussionID = discussionID   // Устанавливаем ID обсуждения для каждого сообщения
		err := r.CreateMessage(ctx, &message) // Добавляем сообщение к обсуждению
		if err != nil {
			// Обработка ошибки, если не удалось добавить сообщение
			return err
		}
	}

	return nil
}

// LeaveDiscussion удаляет пользователя из обсуждения в базе данных
func (r *DiscussionRepository) LeaveDiscussion(userID, discussionID int) error {
	query := "DELETE FROM user_discussions WHERE user_id = $1 AND discussion_id = $2"
	_, err := r.DB.Exec(query, userID, discussionID)
	if err != nil {
		return err
	}
	return nil
}

// getUsersMessages - пример функции, которая возвращает сообщения пользователя
func getUsersMessages(userID int) []models.Message {
	// Здесь можно реализовать логику для получения сообщений пользователя из другого источника, например, из другой таблицы или сервиса
	return []models.Message{
		{UserID: userID, Text: "Привет, мир!"},
		{UserID: userID, Text: "Это сообщение от нового пользователя."},
	}
}

// CreateMessage создает новое сообщение в указанном обсуждении.
func (r *DiscussionRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	query := `INSERT INTO messages (discussion_id, user_id, text, created_at)
			  VALUES ($1, $2, $3, $4)`

	_, err := r.DB.ExecContext(ctx, query, message.DiscussionID, message.UserID, message.Text, message.CreationTime)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMessage обновляет существующее сообщение в указанном обсуждении.
func (r *DiscussionRepository) UpdateMessage(ctx context.Context, message *models.Message) error {
	query := `UPDATE messages
			  SET text = $1, updated_at = $2
			  WHERE id = $3`

	_, err := r.DB.ExecContext(ctx, query, message.Text, time.Now(), message.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteMessage удаляет сообщение из указанного обсуждения.
func (r *DiscussionRepository) DeleteMessage(ctx context.Context, messageID int) error {
	query := `DELETE FROM messages WHERE id = $1`

	_, err := r.DB.ExecContext(ctx, query, messageID)
	if err != nil {
		return err
	}

	return nil
}

// GetMessageByID возвращает сообщение по его идентификатору.
func (r *DiscussionRepository) GetMessageByID(ctx context.Context, messageID int) (*models.Message, error) {
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

// GetMessagesByDiscussion возвращает все сообщения в указанном обсуждении, включая сообщения всех пользователей, присоединившихся к обсуждению.
func (r *DiscussionRepository) GetMessagesByDiscussion(ctx context.Context, userID, discussionID int) ([]*models.Message, error) {
	query := `SELECT m.id, m.discussion_id, m.user_id, m.text, m.created_at, m.updated_at
			  FROM messages m
			  JOIN user_discussions ud ON m.discussion_id = ud.discussion_id
			  WHERE m.discussion_id = $1 AND ud.user_id = $2`

	rows, err := r.DB.QueryContext(ctx, query, discussionID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.DiscussionID, &message.UserID, &message.Text, &message.CreationTime, &message.LastEditTime); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
