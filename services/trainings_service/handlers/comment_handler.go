package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mindmentor/services/trainings_service/models"
	"mindmentor/services/trainings_service/services"
	"net/http"
	"strconv"
	"time"
)

// CommentHandler представляет обработчик HTTP-запросов для работы с комментариями
type CommentHandler struct {
	CommentServ *services.CommentService
}

func NewCommentHandler(commentServ *services.CommentService) *CommentHandler {
	return &CommentHandler{CommentServ: commentServ}
}

// AddCommentHandler обрабатывает запрос на добавление нового комментария к тренировке
func (h *CommentHandler) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Error decoding comment JSON", err)
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
	err = h.CommentServ.AddComment(trainingID, comment.UserID, comment.Text)
	if err != nil {
		log.Println("Ошибка добавления комментария:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Комментарий успешно добавлен")
}

// GetCommentsByTrainingIDHandler обрабатывает запрос на получение комментариев для указанной тренировки
func (h *CommentHandler) GetCommentsByTrainingIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора тренировки из параметров запроса
	trainingIDStr := r.URL.Query().Get("training_id")
	trainingID, err := strconv.Atoi(trainingIDStr)
	if err != nil {
		log.Println("Не удалось получить идентификатор тренировки", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Получение комментариев для указанной тренировки
	comments, err := h.CommentServ.GetCommentsByTrainingID(trainingID)
	if err != nil {
		log.Println("Ошибка при получении комментариев для тренировки", err)
		http.Error(w, "Bad Request", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(comments)
	if err != nil {
		log.Println("Error marshalling comments to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправка комментариев в виде JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
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
	var rating models.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		log.Println("Error decoding rating JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

// GetRatingHandler обрабатывает запрос на получение среднего рейтинга тренировки
func (h *CommentHandler) GetRatingHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора тренировки из запроса
	trainingID, err := strconv.Atoi(r.URL.Query().Get("training_id"))
	if err != nil {
		http.Error(w, "Некорректный идентификатор тренировки", http.StatusBadRequest)
		return
	}

	// Получение рейтинга тренировки из репозитория
	averageRating, err := h.CommentServ.GetRating(trainingID)
	if err != nil {
		http.Error(w, "Ошибка при получении рейтинга тренировки", http.StatusInternalServerError)
		return
	}

	// Отправка ответа с средним рейтингом в формате JSON
	response := map[string]float64{"average_rating": averageRating}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Ошибка при кодировании ответа:", err)
	}
}
