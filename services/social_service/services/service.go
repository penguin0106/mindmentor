package services

import (
	"mindmentor/services/social_service/repositories"
	"mindmentor/shared/models"
)

type DiscussionService struct {
	Repository *repositories.DiscussionRepository
}

func NewDiscussionService(discussionRepo *repositories.DiscussionRepository) *DiscussionService {
	return &DiscussionService{Repository: discussionRepo}
}

// CreateDiscussion создает новое обсуждение
func (s *DiscussionService) CreateDiscussion(discussion *models.Discussion) error {
	return s.Repository.CreateDiscussion(discussion)
}

// FindDiscussion ищет обсуждение по его теме
func (s *DiscussionService) FindDiscussion(topic string) (*models.Discussion, error) {
	return s.Repository.FindDiscussion(topic)
}

// JoinDiscussion добавляет пользователя к обсуждению
func (s *DiscussionService) JoinDiscussion(userID, discussionID int) error {
	return s.Repository.JoinDiscussion(userID, discussionID)
}

// LeaveDiscussion удаляет пользователя из обсуждения
func (s *DiscussionService) LeaveDiscussion(userID, discussionID int) error {
	return s.Repository.LeaveDiscussion(userID, discussionID)
}
