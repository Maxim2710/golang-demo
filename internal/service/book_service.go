package service

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang-demo/internal/models"
)

type BookService struct {
	DB *sqlx.DB
}

func NewBookService(db *sqlx.DB) *BookService {
	return &BookService{DB: db}
}

func (s *BookService) CreateBook(book *models.Book) error {
	query := `
		INSERT INTO books(title, author) 
		VALUES (:title, :author)
		RETURNING id, created_at
	`
	rows, err := s.DB.NamedQuery(query, book)
	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&book.ID, &book.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BookService) GetBookById(id int64) (*models.Book, error) {
	query := `
		SELECT id, title, author, created_at FROM books WHERE id = $1
	`

	var book models.Book
	err := s.DB.Get(&book, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}

		return nil, err
	}

	return &book, nil
}
