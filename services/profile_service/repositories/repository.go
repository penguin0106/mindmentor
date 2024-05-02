package repositories

import (
	"database/sql"
	"mindmentor/shared/models"
)

type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository создает новый экземпляр репозитория пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(userID int, updatedUser *models.User) error {
	_, err := r.DB.Exec("UPDATE users SET username = $2, email = $3 WHERE id = $1", userID, updatedUser.Username, updatedUser.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetFavoriteCourse(userID int) ([]models.Favorite, error) {
	rows, err := r.DB.Query("SELECT item_id FROM course_favorites WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var itemID int
		err := rows.Scan(&itemID)
		if err != nil {
			return nil, err
		}
		favorite := models.Favorite{
			UserID: userID,
			ItemID: itemID,
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func (r *UserRepository) GetFavoriteTraining(userID int) ([]models.Favorite, error) {
	rows, err := r.DB.Query("SELECT item_id FROM trainings_favorites WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var itemID int
		err := rows.Scan(&itemID)
		if err != nil {
			return nil, err
		}
		favorite := models.Favorite{
			UserID: userID,
			ItemID: itemID,
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}
