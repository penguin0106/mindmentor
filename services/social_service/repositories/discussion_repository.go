package repositories

import (
	"database/sql"
	"errors"
	"mindmentor/shared/models"
)

type DiscussionRepository struct {
	DB *sql.DB
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
func (r *DiscussionRepository) JoinDiscussion(userID, discussionID int) error {
	query := "INSERT INTO user_discussions (user_id, discussion_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, userID, discussionID)
	if err != nil {
		return err
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
