package handlers

import (
	"encoding/json"
	"mindmentor/services/trainings_service/repositories"
	"mindmentor/shared/models"
	"net/http"
)

// TrainingHandler представляет обработчик HTTP-запросов для работы с тренировками
type TrainingHandler struct {
	Repository *repositories.TrainingRepository
}

// GetAllTrainingsHandler обрабатывает запрос на получение всех тренировок
func (h *TrainingHandler) GetAllTrainingsHandler(w http.ResponseWriter, _ *http.Request) {
	// Реализация получения всех тренировок из репозитория
	trainings, err := h.Repository.GetAllTrainings()
	if err != nil {
		http.Error(w, "Ошибка при получении тренировок", http.StatusInternalServerError)
		return
	}

	// Отправка тренировок в виде JSON-ответа
	responseJSON, err := json.Marshal(trainings)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON-ответа", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// AddTrainingHandler обрабатывает запрос на добавление новой тренировки
func (h *TrainingHandler) AddTrainingHandler(w http.ResponseWriter, r *http.Request) {
	var training models.Training
	err := json.NewDecoder(r.Body).Decode(&training)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Реализация добавления новой тренировки в репозиторий
	err = h.Repository.AddTraining(&training)
	if err != nil {
		http.Error(w, "Ошибка при добавлении тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(training)
}
