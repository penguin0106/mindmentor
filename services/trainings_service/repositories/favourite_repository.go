package repositories

import (
	"database/sql"
	"errors"
)

// FavoriteRepository представляет репозиторий для работы с избранными тренировками пользователя
type FavoriteRepository struct {
	DB *sql.DB
}

// AddToFavorites добавляет тренировку в избранное для указанного пользователя
func (r *FavoriteRepository) AddToFavorites(userID, trainingID int) error {
	query := "INSERT INTO favorites (user_id, training_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, userID, trainingID)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return errors.New("ошибка при добавлении тренировки в избранное")
	}

	return nil
}

// RemoveFromFavorites удаляет тренировку из избранного для указанного пользователя
func (r *FavoriteRepository) RemoveFromFavorites(userID, trainingID int) error {
	query := "DELETE FROM favorites WHERE user_id = $1 AND training_id = $2"
	result, err := r.DB.Exec(query, userID, trainingID)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return errors.New("ошибка при удалении тренировки из избранного")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("ошибка при проверке количества удаленных строк")
	}

	// Проверяем, была ли удалена как минимум одна строка
	if rowsAffected == 0 {
		return errors.New("тренировка не найдена в избранном пользователя")
	}

	return nil
}
