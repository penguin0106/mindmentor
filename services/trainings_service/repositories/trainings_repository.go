package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"trainings_service/models"
)

// TrainingRepository представляет репозиторий для работы с тренировками
type TrainingRepository struct {
	DB *sql.DB // Подключение к базе данных
}

func NewTrainingRepository(db *sql.DB) *TrainingRepository {
	return &TrainingRepository{DB: db}
}

// AddBook добавляет новую книгу
func (r *TrainingRepository) AddBook(book *models.Book) error {
	_, err := r.DB.Exec("INSERT INTO books (title, description, content) VALUES ($1, $2, $3)", book.Title, book.Description, book.Content)
	return err
}

// GetBookByID возвращает книгу по ее идентификатору
func (r *TrainingRepository) GetBookByID(bookID int) (*models.Book, error) {
	row := r.DB.QueryRow("SELECT id, title, description, content, created_at FROM books WHERE id = $1", bookID)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Content, &book.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

// GetAllBook возвращает все книги
func (r *TrainingRepository) GetAllBook() ([]*models.Book, error) {
	query := "SELECT id, title, description, content, created_at FROM books"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Content, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookByName возвращает книгу по ее названию
func (r *TrainingRepository) GetBookByName(title string) (*models.Book, error) {
	query := "SELECT id, title, description, content, created_at FROM books WHERE title = $1"
	row := r.DB.QueryRow(query, title)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Content, &book.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}
