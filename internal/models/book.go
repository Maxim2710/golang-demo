package models

import "time"

type Book struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Author    string    `db:"author" json:"author"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
