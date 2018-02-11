package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/pkg/errors"
)

var ReadDB *sql.DB

// ErrInvalidDBProvider is returned in the event that unitialized
// db is used to perform action against.
var ErrInvalidDBProvider = errors.New("invalid DB provider")

// NewMysql creates a new mysql connection.
func NewMysql(dataSource string) error {
	log.Println("Preparing mysql connection...")

	// Open a database value. Specify the mysql driver
	// for database/sql
	var err error
	ReadDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		return errors.Wrapf(err, "[sql.Open] %s", dataSource)
	}

	// Set the connection timeout.
	timeout := 60 * time.Second

	ReadDB.SetMaxIdleConns(200)
	ReadDB.SetConnMaxLifetime(timeout)

	// sql.Open() does not establish any connection to the
	// database. It just prepare the database connection value
	// for later use. To make sure the database is available and
	// accessible, we will use db.Ping()
	if err = ReadDB.Ping(); err != nil {
		return errors.Wrap(err, "[db.Ping]")
	}
	log.Println("Able to connect to the mysql database!")

	return nil
}

// // Cruder provides support for crud operations.
// type Cruder interface {
// 	ListBooks() ([]byte, error)
// 	RetrieveBook(id string) ([]byte, error)
// 	UpdateBook(b *Book) ([]byte, error)
// 	CreateBook(b *Book) ([]byte, error)
// 	DeleteBook(id string) error
// 	ListAuthors() ([]byte, error)
// 	RetrieveAuthor(id string) ([]byte, error)
// 	UpdateAuthor(a *Author) ([]byte, error)
// 	CreateAuthor(a *Author) ([]byte, error)
// 	DeleteAuthor(id string) error
// }

// type DB struct {
// 	*sql.DB
// }

// // NewMysql creates a new mysql connection.
// func NewMysql(dataSource string) (*DB, error) {
// 	db, err := sql.Open("mysql", dataSource)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, ": sql.Open : %s", dataSource)
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, errors.Wrapf(err, ": db.Ping :")
// 	}

// 	return &DB{db}, nil
// }
