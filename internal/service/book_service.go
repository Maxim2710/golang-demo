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
		SELECT id, title, author, created_at 
		FROM books 
		WHERE id = $1
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

func (s *BookService) DeleteBookById(id int64) error {
	query := `
		DELETE FROM books 
		WHERE id = $1
	`

	result, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

func (s *BookService) UpdateBookById(book *models.Book) (*models.Book, error) {
	query := `
		UPDATE books 
		SET title = COALESCE(NULLIF($1, ''), title),
		    author = COALESCE(NULLIF($2, ''), author)
		WHERE id = $3
		RETURNING *
		`

	var updatedBook models.Book

	err := s.DB.QueryRowx(query, book.Title, book.Author, book.ID).StructScan(&updatedBook)

	if err != nil {
		return nil, err
	}

	return &updatedBook, nil
}
