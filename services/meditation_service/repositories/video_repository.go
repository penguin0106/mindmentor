package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"meditation_service/models"
)

type VideoRepository struct {
	DB *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{DB: db}
}

// AddVideo добавляет новый видеофайл курса медитации
func (r *VideoRepository) AddVideo(video *models.Video) error {
	_, err := r.DB.Exec("INSERT INTO meditation_videos (title, description, video_content) VALUES ($1, $2, $3)", video.Title, video.Description, video.VideoContent)
	return err
}

// GetVideoByID возвращает видеофайл курса медитации по его идентификатору
func (r *VideoRepository) GetVideoByID(videoID int) (*models.Video, error) {
	row := r.DB.QueryRow("SELECT id, title, description, video_content, created_at FROM meditation_videos WHERE id = $1", videoID)

	var video models.Video
	err := row.Scan(&video.ID, &video.Title, &video.Description, &video.VideoContent, &video.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &video, nil
}

// GetAllVideos возвращает все видеофайлы
func (r *VideoRepository) GetAllVideos() ([]*models.Video, error) {
	rows, err := r.DB.Query("SELECT id, title, description, created_at FROM meditation_videos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*models.Video
	for rows.Next() {
		var video models.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.CreatedAt); err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *VideoRepository) DeleteVideo(videoID int) error {
	query := "DELETE FROM meditation_videos WHERE id = $1"
	result, err := r.DB.Exec(query, videoID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("видеофайл с указанным ID не найден")
	}
	return nil
}

func (r *VideoRepository) GetVideoByTitle(title string) (*models.Video, error) {
	row := r.DB.QueryRow("SELECT id, title, description, created_at, video_content FROM meditation_videos WHERE title = $1", title)

	var video models.Video
	err := row.Scan(&video.ID, &video.Title, &video.Description, &video.VideoContent, &video.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &video, nil
}
