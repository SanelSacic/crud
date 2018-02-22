package db

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

var ReadDB *sql.DB

// ErrInvalidDBProvider is returned in the event that unitialized
// db is used to perform action against.
var ErrInvalidDBProvider = errors.New("invalid DB provider")

// NewMysql creates a new mysql connection.
func NewMySql() error {
	log.Println("Preparing mysql connection...")

	// Open a database value. Specify the mysql driver
	// for database/sql
	var err error
	ReadDB, err = sql.Open("mysql", DataSource)
	if err != nil {
		return errors.Wrapf(err, "[sql.Open] %s", DataSource)
	}

	// Defined in config.go
	ReadDB.SetMaxIdleConns(MaxIdleConnections)
	ReadDB.SetConnMaxLifetime(TimeoutConnection)

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
