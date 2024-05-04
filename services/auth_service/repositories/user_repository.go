package repositories

import (
	"database/sql"
	"mindmentor/services/auth_service/models"
)

// UserRepository представляет репозиторий пользователей
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository создает новый экземпляр репозитория пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Save сохраняет пользователя в базе данных
func (repo *UserRepository) Save(user *models.User) error {

	_, err := repo.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail ищет пользователя по его email
func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {

	var user models.User
	err := repo.DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Пользователь с таким email не найден
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CheckUserExists(userID string) (bool, error) {
	var exists bool
	err := repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
