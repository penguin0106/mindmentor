package services

import (
	"context"
	"errors"
	"social_service/models"
	"social_service/repositories"
)

type DiscussionService struct {
	Repo *repositories.DiscussionRepository
}

func NewDiscussionService(repo *repositories.DiscussionRepository) *DiscussionService {
	return &DiscussionService{Repo: repo}
}

// AddDiscussion создает новый чат (обсуждение)
func (s *DiscussionService) AddDiscussion(ctx context.Context, topic string, ownerID int) (*models.Discussion, error) {
	// Валидация данных (например, проверка наличия владельца)
	if ownerID == 0 {
		return nil, errors.New("ownerID cannot be empty")
	}

	discussion := &models.Discussion{
		Topic:   topic,
		OwnerID: ownerID,
	}

	err := s.Repo.CreateDiscussion(discussion)
	if err != nil {
		return nil, err
	}

	return discussion, nil
}

// FindDiscussionByTopic ищет чат (обсуждение) по его теме
func (s *DiscussionService) FindDiscussionByTopic(ctx context.Context, topic string) (*models.Discussion, error) {
	return s.Repo.FindDiscussionByTopic(topic)
}

// JoinDiscussion добавляет пользователя к чату (обсуждению)
func (s *DiscussionService) JoinDiscussion(ctx context.Context, userID, discussionID int) error {
	// Валидация данных (например, проверка наличия пользователя и чата)
	if userID == 0 || discussionID == 0 {
		return errors.New("userID or discussionID cannot be empty")
	}

	return s.Repo.JoinDiscussion(userID, discussionID)
}

// LeaveDiscussion удаляет пользователя из чата (обсуждения)
func (s *DiscussionService) LeaveDiscussion(ctx context.Context, userID, discussionID int) error {
	// Валидация данных (например, проверка наличия пользователя и чата)
	if userID == 0 || discussionID == 0 {
		return errors.New("userID or discussionID cannot be empty")
	}

	return s.Repo.LeaveDiscussion(userID, discussionID)
}
