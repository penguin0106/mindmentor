package repositories

import (
	"database/sql"
	"errors"
)

// FavoriteRepository представляет собой репозиторий для работы с избранными элементами
type FavoriteRepository struct {
	DB *sql.DB // Поле для подключения к базе данных
}

func NewFavoriterepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{DB: db}
}

// AddToFavorite добавляет элемент в избранное для указанного пользователя
func (r *FavoriteRepository) AddToFavorite(userID, itemID int) error {
	// Проверка наличия элемента в избранном пользователя
	// Если элемент уже существует, вернуть ошибку
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM course_favorites WHERE user_id = $1 AND item_id = $2", userID, itemID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("элемент уже существует в избранном пользователя")
	}

	// Добавление элемента в избранное
	_, err = r.DB.Exec("INSERT INTO course_favorites (user_id, item_id) VALUES ($1, $2)", userID, itemID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveFromFavorite удаляет элемент из избранного для указанного пользователя
func (r *FavoriteRepository) RemoveFromFavorite(userID, itemID int) error {
	// Проверка наличия элемента в избранном пользователя
	// Если элемент не существует, вернуть ошибку
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM course_favorites WHERE user_id = $1 AND item_id = $2", userID, itemID).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("элемент не существует в избранном пользователя")
	}

	// Удаление элемента из избранного
	result, err := r.DB.Exec("DELETE FROM course_favorites WHERE user_id = $1 AND item_id = $2", userID, itemID)
	if err != nil {
		return err
	}

	// Проверка количества удаленных строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("элемент не найден в избранном пользователя")
	}

	return nil
}
