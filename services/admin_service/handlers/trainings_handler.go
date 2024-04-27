package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/admin_service/repositories"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

// TrainingsHandler представляет обработчик запросов для управления тренировками
type TrainingsHandler struct {
	Repository *repositories.TrainingsServiceRepository
}

// AddTrainingHandler обрабатывает запрос на добавление новой тренировки
func (h *TrainingsHandler) AddTrainingHandler(w http.ResponseWriter, r *http.Request) {
	var training models.Training
	err := json.NewDecoder(r.Body).Decode(&training)
	if err != nil {
		http.Error(w, "Ошибка при декодировании тела запроса", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddTraining(training)
	if err != nil {
		http.Error(w, "Ошибка при добавлении тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Тренировка успешно добавлена")
}

// UpdateTrainingHandler обрабатывает запрос на обновление информации о тренировке
func (h *TrainingsHandler) UpdateTrainingHandler(w http.ResponseWriter, r *http.Request) {
	trainingID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID тренировки должен быть числом", http.StatusBadRequest)
		return
	}

	var training models.Training
	err = json.NewDecoder(r.Body).Decode(&training)
	if err != nil {
		http.Error(w, "Ошибка при декодировании тела запроса", http.StatusBadRequest)
		return
	}
	training.ID = trainingID

	err = h.Repository.UpdateTraining(training)
	if err != nil {
		http.Error(w, "Ошибка при обновлении тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Информация о тренировке успешно обновлена")
}

// DeleteTrainingHandler обрабатывает запрос на удаление тренировки
func (h *TrainingsHandler) DeleteTrainingHandler(w http.ResponseWriter, r *http.Request) {
	trainingID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID тренировки должен быть числом", http.StatusBadRequest)
		return
	}

	err = h.Repository.DeleteTraining(trainingID)
	if err != nil {
		http.Error(w, "Ошибка при удалении тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Тренировка успешно удалена")
}

// GetTrainingsHandler обрабатывает запрос на получение списка всех тренировок
func (h *TrainingsHandler) GetTrainingsHandler(w http.ResponseWriter, _ *http.Request) {
	trainings, err := h.Repository.GetTrainings()
	if err != nil {
		http.Error(w, "Ошибка при получении списка тренировок", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainings)
}
