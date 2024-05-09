package repositories

import (
	"auth_service/models"
	"database/sql"
	"errors"
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
	// Проверяем уникальность email
	existingUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("пользователь с таким email уже зарегистрирован")
	}

	// Проверяем уникальность username
	existingUsername, err := repo.FindByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUsername != nil {
		return errors.New("данный логин уже занят, выберите другой")
	}

	_, err = repo.DB.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// FindByEmail ищет пользователя по его email
func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Пользователь с таким email не найден
		}
		return nil, err
	}
	return &user, nil
}

// FindByUsername ищет пользователя по его username
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Пользователь с таким username не найден
		}
		return nil, err
	}
	return &user, nil
}

// Authenticate пытается аутентифицировать пользователя по его учетным данным
func (repo *UserRepository) Authenticate(identifier, password string) (*models.User, error) {
	// Проверяем, является ли identifier email
	user, err := repo.FindByEmail(identifier)
	if err != nil {
		return nil, err
	}
	// Если пользователь не найден по email, пытаемся найти по username
	if user == nil {
		user, err = repo.FindByUsername(identifier)
		if err != nil {
			return nil, err
		}
	}
	// Если пользователь найден, проверяем совпадение пароля
	if user != nil && user.Password == password {
		return user, nil
	}
	// Если пользователь не найден или пароль не совпадает, возвращаем ошибку
	return nil, errors.New("неверные учетные данные")
}
