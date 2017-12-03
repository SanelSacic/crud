package models

import (
	"database/sql"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

// Cruder provides support for crud operations.
type Cruder interface {
	ListBooks() ([]byte, error)
	RetrieveBook(id string) ([]byte, error)
	UpdateBook(b *Book) ([]byte, error)
	CreateBook(b *Book) ([]byte, error)
	DeleteBook(id string) error
	ListAuthors() ([]byte, error)
	RetrieveAuthor(id string) ([]byte, error)
	UpdateAuthor(a *Author) ([]byte, error)
	CreateAuthor(a *Author) ([]byte, error)
	DeleteAuthor(id string) error
}

type DB struct {
	*sql.DB
}

// NewMysql creates a new mysql connection.
func NewMysql(dataSource string) (*DB, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, errors.Wrapf(err, ": sql.Open : %s", dataSource)
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(err, ": db.Ping :")
	}

	return &DB{db}, nil
}
