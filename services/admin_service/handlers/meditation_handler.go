package handlers

import (
	"encoding/json"
	"fmt"
	"mindmentor/services/admin_service/repositories"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

// MeditationHandler представляет обработчик HTTP-запросов для работы с медитациями
type MeditationHandler struct {
	Repository *repositories.MeditationServiceRepository
}

// AddCourseHandler обрабатывает запрос на добавление нового курса медитаций
func (h *MeditationHandler) AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddCourse(course)
	if err != nil {
		http.Error(w, "Ошибка при добавлении курса медитаций", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Курс медитаций успешно добавлен")
}

// UpdateCourseHandler обрабатывает запрос на обновление информации о курсе медитаций
func (h *MeditationHandler) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.UpdateCourse(course)
	if err != nil {
		http.Error(w, "Ошибка при обновлении информации о курсе медитаций", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Информация о курсе медитаций успешно обновлена")
}

// DeleteCourseHandler обрабатывает запрос на удаление курса медитаций
func (h *MeditationHandler) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	courseIDStr := r.URL.Query().Get("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		http.Error(w, "Некорректный идентификатор курса медитаций", http.StatusBadRequest)
		return
	}

	err = h.Repository.DeleteCourse(courseID)
	if err != nil {
		http.Error(w, "Ошибка при удалении курса медитаций", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Курс медитаций успешно удален")
}

// GetCoursesHandler обрабатывает запрос на получение списка всех курсов медитаций
func (h *MeditationHandler) GetCoursesHandler(w http.ResponseWriter, _ *http.Request) {
	courses, err := h.Repository.GetCourses()
	if err != nil {
		http.Error(w, "Ошибка при получении списка курсов медитаций", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// AddMusicHandler обрабатывает запрос на добавление новой музыки для медитаций
func (h *MeditationHandler) AddMusicHandler(w http.ResponseWriter, r *http.Request) {
	var music models.Music
	err := json.NewDecoder(r.Body).Decode(&music)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.AddMusic(music)
	if err != nil {
		http.Error(w, "Ошибка при добавлении музыки для медитаций", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Музыка для медитаций успешно добавлена")
}

// UpdateMusicHandler обрабатывает запрос на обновление информации о музыке для медитаций
func (h *MeditationHandler) UpdateMusicHandler(w http.ResponseWriter, r *http.Request) {
	var music models.Music
	err := json.NewDecoder(r.Body).Decode(&music)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err = h.Repository.UpdateMusic(music)
	if err != nil {
		http.Error(w, "Ошибка при обновлении информации о музыке для медитаций", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Информация о музыке для медитаций успешно обновлена")
}

// DeleteMusicHandler обрабатывает запрос на удаление музыки для медитаций
func (h *MeditationHandler) DeleteMusicHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметра "id" из URL
	queryValues := r.URL.Query()
	musicID := queryValues.Get("id")
	id, err := strconv.Atoi(musicID)
	if err != nil {
		http.Error(w, "ID музыки должен быть числом", http.StatusBadRequest)
		return
	}

	// Удаление музыки с указанным ID
	err = h.Repository.DeleteMusic(id)
	if err != nil {
		http.Error(w, "Ошибка при удалении музыки для медитаций", http.StatusInternalServerError)
		return
	}

	// Возвращение успешного ответа
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Музыка для медитаций успешно удалена")
}

// GetMusicHandler обрабатывает запрос на получение списка всей музыки для медитаций
func (h *MeditationHandler) GetMusicHandler(w http.ResponseWriter, _ *http.Request) {
	music, err := h.Repository.GetMusic()
	if err != nil {
		http.Error(w, "Ошибка при получении списка музыки для медитаций", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(music)
}
