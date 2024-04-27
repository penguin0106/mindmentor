package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"mindmentor/services/trainings_service/models"
	"mindmentor/services/trainings_service/repositories"
	"net/http"
	"strconv"
	"time"
)

// CommentHandler представляет обработчик HTTP-запросов для работы с комментариями
type CommentHandler struct {
	Repository *repositories.CommentRepository
}

// AddCommentHandler обрабатывает запрос на добавление нового комментария к тренировке
func (h *CommentHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Проверка наличия необходимых полей в комментарии
	if comment.UserID == 0 || comment.Text == "" {
		http.Error(w, "Недостаточно данных для добавления комментария", http.StatusBadRequest)
		return
	}

	// Установка временной метки комментария
	comment.Timestamp = time.Now().Unix()

	// Получение идентификатора тренировки из параметров запроса
	trainingID, err := getTrainingIDFromRequest(r)
	if err != nil {
		http.Error(w, "Не удалось получить идентификатор тренировки", http.StatusBadRequest)
		return
	}

	// Добавление комментария в хранилище
	err = h.Repository.AddComment(trainingID, comment.UserID, comment.Text)
	if err != nil {
		http.Error(w, "Ошибка добавления комментария", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Комментарий успешно добавлен")
}

// GetCommentsByTrainingIDHandler обрабатывает запрос на получение комментариев для указанной тренировки
func (h *CommentHandler) GetCommentsByTrainingIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора тренировки из параметров запроса
	trainingID, err := getTrainingIDFromRequest(r)
	if err != nil {
		http.Error(w, "Не удалось получить идентификатор тренировки", http.StatusBadRequest)
		return
	}

	// Получение комментариев для указанной тренировки
	comments, err := h.Repository.GetCommentsByTrainingID(trainingID)
	if err != nil {
		http.Error(w, "Ошибка при получении комментариев для тренировки", http.StatusInternalServerError)
		return
	}

	// Отправка комментариев в виде JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// getTrainingIDFromRequest извлекает идентификатор тренировки из запроса
func getTrainingIDFromRequest(r *http.Request) (int, error) {
	// Извлекаем параметр "training_id" из строки запроса
	trainingID := r.URL.Query().Get("training_id")

	// Проверяем, что параметр не пустой
	if trainingID == "" {
		return 0, errors.New("идентификатор тренировки не указан")
	}

	// Преобразуем полученный идентификатор в целочисленное значение
	id, err := strconv.Atoi(trainingID)
	if err != nil {
		// Если произошла ошибка преобразования, возвращаем ошибку
		return 0, err
	}

	// Возвращаем полученный идентификатор тренировки
	return id, nil
}

// AddRatingHandler обрабатывает запрос на добавление оценки тренировки
func (h *CommentHandler) AddRatingHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметры из тела запроса
	var requestData struct {
		TrainingID int     `json:"trainingId"`
		Rating     float64 `json:"rating"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Проверяем, что оба параметра присутствуют
	if requestData.TrainingID == 0 || requestData.Rating == 0 {
		http.Error(w, "Некорректные параметры", http.StatusBadRequest)
		return
	}

	// Вызываем метод репозитория для добавления оценки тренировки
	err = h.Repository.AddRating(requestData.TrainingID, requestData.Rating)
	if err != nil {
		http.Error(w, "Ошибка при добавлении оценки тренировки", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CommentHandler) GetRatingHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора тренировки из запроса
	trainingID, err := getTrainingIDFromRequest(r)
	if err != nil {
		http.Error(w, "Не удалось получить идентификатор тренировки", http.StatusBadRequest)
		return
	}

	// Выполнение запроса к репозиторию для получения рейтинга тренировки
	rating, err := h.Repository.GetRating(trainingID)
	if err != nil {
		http.Error(w, "Ошибка при получении рейтинга тренировки", http.StatusInternalServerError)
		return
	}

	// Отправка рейтинга в виде JSON-ответа
	responseJSON, err := json.Marshal(rating)
	if err != nil {
		http.Error(w, "Ошибка при формировании JSON-ответа", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
