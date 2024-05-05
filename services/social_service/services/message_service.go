package services

import (
	"context"
	"errors"
	"social_service/models"
	"social_service/repositories"
)

type MessageService struct {
	Repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{Repo: repo}
}

// SendMessage отправляет сообщение в указанный чат
func (s *MessageService) SendMessage(ctx context.Context, message *models.Message) error {
	// Валидация данных (например, проверка наличия текста сообщения и идентификатора чата)
	if message == nil || message.Text == "" || message.DiscussionID == 0 {
		return errors.New("message text or discussionID cannot be empty")
	}

	return s.Repo.CreateMessage(ctx, message)
}

// EditMessage изменяет существующее сообщение
func (s *MessageService) EditMessage(ctx context.Context, message *models.Message) error {
	// Валидация данных (например, проверка наличия текста сообщения и идентификатора сообщения)
	if message == nil || message.Text == "" || message.ID == 0 {
		return errors.New("message text or ID cannot be empty")
	}

	return s.Repo.UpdateMessage(ctx, message)
}

// DeleteMessage удаляет сообщение
func (s *MessageService) DeleteMessage(ctx context.Context, messageID int) error {
	// Валидация данных (например, проверка наличия идентификатора сообщения)
	if messageID == 0 {
		return errors.New("messageID cannot be empty")
	}

	return s.Repo.DeleteMessage(ctx, messageID)
}
