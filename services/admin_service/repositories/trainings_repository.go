package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"mindmentor/shared/models"
	"net/http"
	"strconv"
)

// TrainingsServiceRepository представляет репозиторий для взаимодействия с сервисом тренировок
type TrainingsServiceRepository struct {
	BaseURL string
}

// AddTraining добавляет новую тренировку в сервис тренировок
func (r *TrainingsServiceRepository) AddTraining(training models.Training) error {
	requestBody, err := json.Marshal(training)
	if err != nil {
		return err
	}

	_, err = http.Post(r.BaseURL+"/trainings", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	return nil
}

// UpdateTraining обновляет информацию о тренировке в сервисе тренировок
func (r *TrainingsServiceRepository) UpdateTraining(training models.Training) error {
	requestBody, err := json.Marshal(training)
	if err != nil {
		return err
	}

	url := r.BaseURL + "/trainings/" + strconv.Itoa(training.ID)
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
		return errors.New("не удалось обновить информацию о тренировке")
	}

	return nil
}

// DeleteTraining удаляет тренировку из сервиса тренировок
func (r *TrainingsServiceRepository) DeleteTraining(trainingID int) error {
	req, err := http.NewRequest("DELETE", r.BaseURL+"/trainings/"+strconv.Itoa(trainingID), nil)
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

// GetTrainings возвращает список всех тренировок из сервиса тренировок
func (r *TrainingsServiceRepository) GetTrainings() ([]models.Training, error) {
	resp, err := http.Get(r.BaseURL + "/trainings")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var trainings []models.Training
	err = json.NewDecoder(resp.Body).Decode(&trainings)
	if err != nil {
		return nil, err
	}

	return trainings, nil
}
