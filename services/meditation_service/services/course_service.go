package services

import (
	"errors"
	"meditation_service/models"
	"meditation_service/repositories"
)

type CourseService struct {
	CourseRepo *repositories.CourseRepository
}

func NewCourseService(courService *repositories.CourseRepository) *CourseService {
	return &CourseService{CourseRepo: courService}
}

// GetAllCourses возвращает все курсы медитации
func (s *CourseService) GetAllCourses() ([]*models.Course, error) {
	return s.CourseRepo.GetAllCourses()
}

// GetCourseByName возвращает курс медитации по его названию
func (s *CourseService) GetCourseByName(courseName string) (*models.Course, error) {
	course, err := s.CourseRepo.GetCourseByName(courseName)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, errors.New("Курс с указанным названием не найден")
	}
	return course, nil
}

// AddCourse добавляет новый курс медитации
func (s *CourseService) AddCourse(course *models.Course) error {
	return s.CourseRepo.AddCourse(course)
}
