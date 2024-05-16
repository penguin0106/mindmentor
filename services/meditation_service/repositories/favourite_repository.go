package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"meditation_service/models"
)

// FavoriteRepository представляет собой репозиторий для работы с избранными элементами
type FavoriteRepository struct {
	DB *sql.DB // Поле для подключения к базе данных
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{DB: db}
}

// AddToFavorite добавляет элемент в избранное для указанного пользователя
func (r *FavoriteRepository) AddToFavorite(fav *models.Favorite) error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM video_favorites WHERE user_id = $1 AND item_id = $2", fav.UserID, fav.ItemID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("видео уже существует в избранном пользователя")
	}

	_, err = r.DB.Exec("INSERT INTO video_favorites (user_id, item_id) VALUES ($1, $2)", fav.UserID, fav.ItemID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFromFavorite удаляет элемент из избранного для указанного пользователя
func (r *FavoriteRepository) RemoveFromFavorite(fav *models.Favorite) error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM video_favorites WHERE user_id = $1 AND item_id = $2", fav.UserID, fav.ItemID).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("видео не существует в избранном пользователя")
	}

	result, err := r.DB.Exec("DELETE FROM video_favorites WHERE user_id = $1 AND item_id = $2", fav.UserID, fav.ItemID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("видео не найдено в избранном пользователя")
	}
	return nil
}
