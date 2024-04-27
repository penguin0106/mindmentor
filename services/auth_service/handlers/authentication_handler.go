package handlers

import (
	"encoding/json"
	"mindmentor/services/auth_service/repositories"
	"net/http"
)

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем логин и пароль из запроса
	login := r.FormValue("login")
	password := r.FormValue("password")

	// Проводим аутентификацию, проверяем логин и пароль в базе данных
	user, err := repositories.UserRepository.GetUserByLogin(login)
	if err != nil {
		// Если произошла ошибка при получении пользователя из базы данных, возвращаем ошибку сервера
		http.Error(w, "Ошибка аутентификации", http.StatusInternalServerError)
		return
	}

	// Проверяем, совпадает ли пароль пользователя с предоставленным паролем
	if user == nil || user.Password != password {
		// Если пароль не совпадает или пользователь не найден, возвращаем ошибку аутентификации
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	// Если аутентификация успешна, генерируем токен доступа для пользователя
	token, err := repositories.TokenRepository.GenerateToken(user.ID)
	if err != nil {
		// Если произошла ошибка при генерации токена, возвращаем ошибку сервера
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	// Формируем ответ в формате JSON с токеном доступа
	response := map[string]string{
		"token": token,
	}

	// Кодируем ответ в JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Если произошла ошибка при кодировании JSON, возвращаем ошибку сервера
		http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем ответ с токеном доступа
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
