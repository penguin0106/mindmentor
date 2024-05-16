package repositories

import (
	"database/sql"
	"errors"
	"profile_service/models"
)

type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository создает новый экземпляр репозитория пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserByID возвращает данные пользователя по его идентификатору
func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	return &user, nil
}

// EditProfileUsername редактирует имя пользователя в профиле
func (r *UserRepository) EditProfileUsername(userID int, newUsername string) error {
	query := "UPDATE users SET username = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}

// EditProfileEmail редактирует email в профиле
func (r *UserRepository) EditProfileEmail(userID int, newEmail string) error {
	query := "UPDATE users SET email = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, newEmail, userID)
	if err != nil {
		return err
	}
	return nil
}

// EditProfilePassword редактирует пароль в профиле
func (r *UserRepository) EditProfilePassword(userID int, newPassword string) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, newPassword, userID)
	if err != nil {
		return err
	}
	return nil
}

// GetFavoriteVideos возвращает избранные видео пользователя
func (r *UserRepository) GetFavoriteVideos(userID int) ([]int, error) {
	query := "SELECT video_id FROM video_favorites WHERE user_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favoriteVideos []int
	for rows.Next() {
		var fav models.Favorite
		err := rows.Scan(&fav.ItemID)
		if err != nil {
			return nil, err
		}
		favoriteVideos = append(favoriteVideos, fav.ItemID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favoriteVideos, nil
}

// GetFavoriteTrainings возвращает избранные тренировки пользователя
func (r *UserRepository) GetFavoriteTrainings(userID int) ([]int, error) {
	query := "SELECT training_id FROM training_favorites WHERE user_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favoriteTrainings []int
	for rows.Next() {
		var fav models.Favorite
		err := rows.Scan(&fav.ItemID)
		if err != nil {
			return nil, err
		}
		favoriteTrainings = append(favoriteTrainings, fav.ItemID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favoriteTrainings, nil
}
