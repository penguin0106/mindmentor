package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mindmentor/activity_tracker_service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Установка обработчика для эндпоинта добавления данных об активности
	router.HandleFunc("/activity", handlers.AddActivity).Methods("POST")

	log.Fatal(http.ListenAndServe(":8085", router))
}
