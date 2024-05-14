package services

import (
	"meditation_service/models"
	"meditation_service/repositories"
)

type VideoService struct {
	VideoRepo *repositories.VideoRepository
}

func NewVideoService(videoRepo *repositories.VideoRepository) *VideoService {
	return &VideoService{VideoRepo: videoRepo}
}

func (s *VideoService) GetAllVideos() ([]*models.Video, error) {
	videos, err := s.VideoRepo.GetAllVideos()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (s *VideoService) AddVideo(title, description string, videoContent []byte) error {
	video := &models.Video{
		Title:        title,
		Description:  description,
		VideoContent: videoContent,
	}

	err := s.VideoRepo.AddVideo(video)
	if err != nil {
		return err
	}
	return nil
}

func (s *VideoService) GetVideoByID(videoID int) (*models.Video, error) {
	video, err := s.VideoRepo.GetVideoByID(videoID)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (s *VideoService) DeleteVideo(videoID int) error {
	err := s.VideoRepo.DeleteVideo(videoID)
	if err != nil {
		return err
	}
	return nil
}

func (s *VideoService) GetVideoByTitle(title string) (*models.Video, error) {
	video, err := s.VideoRepo.GetVideoByTitle(title)
	if err != nil {
		return nil, err
	}
	return video, nil
}
