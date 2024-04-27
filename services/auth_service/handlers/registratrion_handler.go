package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/auth_service/repositories"
	"mindmentor/shared/models"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем данные нового пользователя из запроса
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	// Проводим проверки валидности данных пользователя
	if user.Login == "" || user.Password == "" {
		http.Error(w, "Логин и пароль обязательны", http.StatusBadRequest)
		return
	}

	// Регистрируем нового пользователя в базе данных
	err = repositories.UserRepository.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Пользователь успешно зарегистрирован")
}
