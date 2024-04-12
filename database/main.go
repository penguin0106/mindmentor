package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"mindmentor/database/handlers"
	"net/http"
	"os"
)

func main() {
	// Получаем параметры подключения к базе данных из переменных окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Формируем строку подключения к базе данных
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal("Error checking database connection: ", err)
	}
	log.Println("Connected to database")

	// Инициализируем объект базы данных в пакете handlers
	handlers.InitDB(db)

	// Инициализируем маршрутизатор
	router := mux.NewRouter()

	// Устанавливаем обработчики запросов
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Запускаем HTTP сервер на порту 8080
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
