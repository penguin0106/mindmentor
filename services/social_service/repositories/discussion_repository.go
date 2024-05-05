package repositories

import (
	"database/sql"
	"errors"
	"social_service/models"
)

type DiscussionRepository struct {
	DB *sql.DB
}

func NewDiscussionRepository(db *sql.DB) *DiscussionRepository {
	return &DiscussionRepository{DB: db}
}

// CreateDiscussion создает новый чат (обсуждение) в базе данных
func (r *DiscussionRepository) CreateDiscussion(discussion *models.Discussion) error {
	query := "INSERT INTO discussions (topic, owner_id) VALUES ($1, $2) RETURNING id"
	err := r.DB.QueryRow(query, discussion.Topic, discussion.OwnerID).Scan(&discussion.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindDiscussionByTopic ищет чат по его теме в базе данных
func (r *DiscussionRepository) FindDiscussionByTopic(topic string) (*models.Discussion, error) {
	query := "SELECT id, topic, owner_id FROM discussions WHERE topic = $1"
	row := r.DB.QueryRow(query, topic)

	var discussion models.Discussion
	err := row.Scan(&discussion.ID, &discussion.Topic, &discussion.OwnerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Чат с указанной темой не найден
			return nil, nil
		}
		return nil, err
	}

	return &discussion, nil
}

// JoinDiscussion добавляет пользователя к чату (обсуждению) в базе данных.
func (r *DiscussionRepository) JoinDiscussion(userID, discussionID int) error {
	query := "INSERT INTO user_discussions (user_id, discussion_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, userID, discussionID)
	if err != nil {
		return err
	}
	return nil
}

// LeaveDiscussion удаляет пользователя из чата в базе данных
func (r *DiscussionRepository) LeaveDiscussion(userID, discussionID int) error {
	query := "DELETE FROM user_discussions WHERE user_id = $1 AND discussion_id = $2"
	_, err := r.DB.Exec(query, userID, discussionID)
	if err != nil {
		return err
	}
	return nil
}
