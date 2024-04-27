package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

// MeditationServiceRepository представляет репозиторий для взаимодействия с сервисом медитаций
type MeditationServiceRepository struct {
	BaseURL string
}

// AddCourse добавляет новый курс медитаций в сервис медитаций
func (r *MeditationServiceRepository) AddCourse(course models.Course) error {
	// Преобразование структуры в JSON
	requestBody, err := json.Marshal(course)
	if err != nil {
		return err
	}

	// Отправка POST запроса на добавление курса медитаций
	_, err = http.Post(r.BaseURL+"/courses", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	return nil
}

// UpdateCourse обновляет информацию о курсе медитаций в сервисе медитаций
func (r *MeditationServiceRepository) UpdateCourse(course models.Course) error {
	// Преобразование структуры в JSON
	requestBody, err := json.Marshal(course)
	if err != nil {
		return err
	}

	// Создание запроса PUT
	req, err := http.NewRequest(http.MethodPut, r.BaseURL+"/courses/"+strconv.Itoa(course.ID), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return errors.New("не удалось обновить информацию о курсе медитаций")
	}

	return nil
}

// DeleteCourse удаляет курс медитаций из сервиса медитаций
func (r *MeditationServiceRepository) DeleteCourse(courseID int) error {
	// Отправка DELETE запроса на удаление курса медитаций
	req, err := http.NewRequest("DELETE", r.BaseURL+"/courses/"+strconv.Itoa(courseID), nil)
	if err != nil {
		return err
	}

	// Отправка запроса
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// GetCourses возвращает список всех курсов медитаций из сервиса медитаций
func (r *MeditationServiceRepository) GetCourses() ([]models.Course, error) {
	// Отправка GET запроса для получения списка всех курсов медитаций
	resp, err := http.Get(r.BaseURL + "/courses")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Декодирование JSON ответа в структуру
	var courses []models.Course
	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

// AddMusic добавляет новую музыку для медитаций в сервис медитаций
func (r *MeditationServiceRepository) AddMusic(music models.Music) error {
	requestBody, err := json.Marshal(music)
	if err != nil {
		return err
	}

	_, err = http.Post(r.BaseURL+"/music", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	return nil
}

// UpdateMusic обновляет информацию о музыке для медитаций в сервисе медитаций
func (r *MeditationServiceRepository) UpdateMusic(music models.Music) error {
	requestBody, err := json.Marshal(music)
	if err != nil {
		return err
	}

	url := r.BaseURL + "/music/" + strconv.Itoa(music.ID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("не удалось обновить музыку")
	}

	return nil
}

// DeleteMusic удаляет музыку для медитаций из сервиса медитаций
func (r *MeditationServiceRepository) DeleteMusic(musicID int) error {
	req, err := http.NewRequest("DELETE", r.BaseURL+"/music/"+strconv.Itoa(musicID), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// GetMusic возвращает список всей музыки для медитаций из сервиса медитаций
func (r *MeditationServiceRepository) GetMusic() ([]models.Music, error) {
	resp, err := http.Get(r.BaseURL + "/music")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var music []models.Music
	err = json.NewDecoder(resp.Body).Decode(&music)
	if err != nil {
		return nil, err
	}

	return music, nil
}
