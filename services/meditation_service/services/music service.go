package services

import (
	"fmt"
	"meditation_service/models"
	"meditation_service/repositories"
)

type MusicService struct {
	MusicRepo *repositories.MusicRepository
}

func NewMusicService(musicService *repositories.MusicRepository) *MusicService {
	return &MusicService{MusicRepo: musicService}
}

// AddMusic добавляет новый аудиофайл для медитации
func (s *MusicService) AddMusic(name string, duration int, musicFile []byte) error {
	music := &models.Music{
		Name:      name,
		Duration:  duration,
		MusicFile: musicFile,
	}

	err := s.MusicRepo.AddMusic(music)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении аудиофайла: %v", err)
	}

	return nil
}

// GetMusicByID возвращает аудиофайл для медитации по его идентификатору
func (s *MusicService) GetMusicByID(musicID int) (*models.Music, error) {
	music, err := s.MusicRepo.GetMusicByID(musicID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аудиофайла по ID: %v", err)
	}
	return music, nil
}

// GetAllMusic возвращает все аудиофайлы для медитации
func (s *MusicService) GetAllMusic() ([]*models.Music, error) {
	musicList, err := s.MusicRepo.GetAllMusic()
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка аудиофайлов: %v", err)
	}
	return musicList, nil
}
