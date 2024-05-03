package services

import (
	"mindmentor/services/meditation_service/models"
	"mindmentor/services/meditation_service/repositories"
)

type MusicService struct {
	MusicRepo *repositories.MusicRepository
}

func NewMusicService(musicService *repositories.MusicRepository) *MusicService {
	return &MusicService{MusicRepo: musicService}
}

// GetAllMusic возвращает все аудиофайлы для медитации
func (s *MusicService) GetAllMusic() ([]*models.Music, error) {
	return s.MusicRepo.GetAllMusic()
}

// GetMusicByID возвращает аудиофайл для медитации по его идентификатору
func (s *MusicService) GetMusicByID(musicID int) (*models.Music, error) {
	return s.MusicRepo.GetMusicByID(musicID)
}

// AddMusic добавляет новый аудиофайл для медитации
func (s *MusicService) AddMusic(music *models.Music) error {
	return s.MusicRepo.AddMusic(music)
}
