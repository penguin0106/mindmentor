package repositories

import (
	"database/sql"
	"errors"
	"mindmentor/services/auth_service/models"
)

// UserRepository представляет собой репозиторий для работы с пользователями
type UserRepository struct {
	DB *sql.DB
}

// GetUserByID получает пользователя из базы данных по его идентификатору
func (r *UserRepository) GetUserByID(userID int) (*models.User, error) {
	query := "SELECT id, login, password FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByLogin получает пользователя из базы данных по его логину
func (r *UserRepository) GetUserByLogin(login string) (*models.User, error) {
	query := "SELECT id, login, password FROM users WHERE login = $1"
	row := r.DB.QueryRow(query, login)

	var user models.User
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser регистрирует нового пользователя в базе данных
func (r *UserRepository) CreateUser(user *models.User) error {
	// Проверяем, существует ли пользователь с таким же логином
	existingUser, err := r.GetUserByLogin(user.Login)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("пользователь с таким логином уже существует")
	}

	// Вставляем нового пользователя в базу данных
	query := "INSERT INTO users (login, password) VALUES ($1, $2)"
	_, err = r.DB.Exec(query, user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}
