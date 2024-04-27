package main

import (
	"fmt"
	"mindmentor/services/auth_service/handlers"
	"net/http"
)

func main() {
	// Инициализация сервера и других компонентов
	fmt.Println("Сервис аутентификации запущен")

	// Запуск HTTP сервера
	http.HandleFunc("/login", handlers.AuthenticationHandler)
	http.HandleFunc("/authorize", handlers.AuthorizationHandler)
	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", nil)
}
