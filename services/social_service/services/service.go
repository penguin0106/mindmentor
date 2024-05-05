package services

import (
	"context"
	"errors"
	"social_service/models"
	"social_service/repositories"
	"time"
)

type DiscussionService struct {
	Repo *repositories.DiscussionRepository
}

func NewDiscussionService(repo *repositories.DiscussionRepository) *DiscussionService {
	return &DiscussionService{Repo: repo}
}

func (s *DiscussionService) CreateDiscussion(discussion *models.Discussion) error {
	return s.Repo.CreateDiscussion(discussion)
}

func (s *DiscussionService) FindDiscussion(topic string) (*models.Discussion, error) {
	return s.Repo.FindDiscussion(topic)
}

func (s *DiscussionService) JoinDiscussion(ctx context.Context, userID, discussionID int, userMessageRepo repositories.UserMessageRepository) error {
	return s.Repo.JoinDiscussion(ctx, userID, discussionID, userMessageRepo)
}

func (s *DiscussionService) LeaveDiscussion(userID, discussionID int) error {
	return s.Repo.LeaveDiscussion(userID, discussionID)
}

func (s *DiscussionService) CreateMessage(ctx context.Context, message *models.Message) error {
	// Установка времени создания сообщения
	message.CreationTime = time.Now()
	return s.Repo.CreateMessage(ctx, message)
}

func (s *DiscussionService) UpdateMessage(ctx context.Context, message *models.Message) error {
	// Проверка существования сообщения
	existingMessage, err := s.Repo.GetMessageByID(ctx, message.ID)
	if err != nil {
		return err
	}
	if existingMessage == nil {
		return errors.New("message not found")
	}

	// Установка времени последнего редактирования сообщения
	message.LastEditTime = time.Now()
	return s.Repo.UpdateMessage(ctx, message)
}

func (s *DiscussionService) DeleteMessage(ctx context.Context, messageID int) error {
	// Проверка существования сообщения
	existingMessage, err := s.Repo.GetMessageByID(ctx, messageID)
	if err != nil {
		return err
	}
	if existingMessage == nil {
		return errors.New("message not found")
	}

	return s.Repo.DeleteMessage(ctx, messageID)
}

func (s *DiscussionService) GetMessagesByDiscussion(ctx context.Context, userID, discussionID int) ([]*models.Message, error) {
	return s.Repo.GetMessagesByDiscussion(ctx, userID, discussionID)
}
