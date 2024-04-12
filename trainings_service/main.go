package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mindmentor/trainings_service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Установка обработчиков для эндпоинтов поиска тренировок и добавления отзывов
	router.HandleFunc("/trainings", handlers.SearchTrainings).Methods("GET")
	router.HandleFunc("/review", handlers.AddReview).Methods("POST")

	log.Fatal(http.ListenAndServe(":8083", router))
}
