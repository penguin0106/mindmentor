package repositories

import (
	"database/sql"
	"errors"
	"meditation_service/models"
)

// CourseRepository представляет собой репозиторий для работы с курсами медитации
type CourseRepository struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

// GetAllCourses возвращает все курсы медитации
func (r *CourseRepository) GetAllCourses() ([]*models.Course, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*models.Course
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

// GetCourseByName возвращает курс медитации по его названию
func (r *CourseRepository) GetCourseByName(courseName string) (*models.Course, error) {
	row := r.DB.QueryRow("SELECT id, name, description FROM courses WHERE name = $1", courseName)

	var course models.Course
	err := row.Scan(&course.ID, &course.Title, &course.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Курс с таким названием не найден
		}
		return nil, err
	}

	return &course, nil
}

// AddCourse добавляет новый курс медитации
func (r *CourseRepository) AddCourse(course *models.Course) error {
	_, err := r.DB.Exec("INSERT INTO courses (name, description) VALUES ($1, $2)", course.Title, course.Description)
	return err
}
