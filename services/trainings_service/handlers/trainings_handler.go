package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"trainings_service/models"
	"trainings_service/services"
)

// TrainingHandler представляет обработчик HTTP-запросов для работы с тренировками
type TrainingHandler struct {
	TrainingServ *services.TrainingService
}

func NewTrainingHandler(trainingServ *services.TrainingService) *TrainingHandler {
	return &TrainingHandler{TrainingServ: trainingServ}
}

// GetAllTrainingsHandler обрабатывает запрос на получение всех тренировок
func (h *TrainingHandler) GetAllTrainingsHandler(w http.ResponseWriter, _ *http.Request) {
	// Реализация получения всех тренировок из репозитория
	trainings, err := h.TrainingServ.GetAllTrainings()
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
	err = h.TrainingServ.AddTraining(&training)
	if err != nil {
		http.Error(w, "Ошибка при добавлении тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(training)
}

func (h *TrainingHandler) GetTrainingByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра из URL
	trainingName := r.URL.Query().Get("name")
	if trainingName == "" {
		http.Error(w, "Missing training name parameter", http.StatusBadRequest)
		return
	}

	// Вызываем соответствующий метод сервиса для получения тренировки по имени
	training, err := h.TrainingServ.GetTrainingByName(trainingName)
	if err != nil {
		log.Println("Error getting training by name:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Если тренировка не найдена, возвращаем соответствующий код ответа
	if training == nil {
		http.Error(w, "Training not found", http.StatusNotFound)
		return
	}

	// Преобразуем найденную тренировку в формат JSON
	response, err := json.Marshal(training)
	if err != nil {
		log.Println("Error marshalling training to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type и код ответа 200 (OK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Отправляем ответ с данными тренировки в формате JSON
	w.Write(response)
}
