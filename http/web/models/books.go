package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Book contains information about the book.
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    string `json:"author-id"`
	Image       string `json:"image"`
	Published   string `json:"published"`
}

// ListBooks retrieves a list of existing books in the database.
func (db *DB) ListBooks() ([]byte, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, errors.Wrap(err, ": conn.Query :")
	}
	defer rows.Close()

	var books []*Book

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.Published, &b.Image, &b.AuthorID)
		if err != nil {
			return nil, errors.Wrapf(err, ": rows.Next: %v", b)
		}
		books = append(books, &b)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, ": rows.Err :")
	}

	var data []byte

	if data, err = json.Marshal(&books); err != nil {
		return nil, errors.Wrap(err, ": json.Marhal :")
	}

	return data, nil
}

// RetrieveBook gets the specific book from the database.
func (db *DB) RetrieveBook(id string) ([]byte, error) {

	rows, err := db.Query("SELECT * FROM books WHERE id = " + id)
	if err != nil {
		return nil, errors.Wrapf(err, ": conn.Query.Retrieve(%s)", rows)
	}
	defer rows.Close()

	var book []*Book

	if rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.Published, &b.Image, &b.AuthorID)
		if err != nil {
			return nil, errors.Wrapf(err, ": rows.Scann : %v", b)
		}
		book = append(book, &b)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, ": rows.Err :")
	}

	var data []byte

	if data, err = json.Marshal(&book); err != nil {
		return nil, errors.Wrapf(err, ": json.Marshal : %d", data)
	}

	return data, nil
}

// CreateBook creates a new Book in the database.
func (db *DB) CreateBook(b *Book) ([]byte, error) {

	stmt, err := db.Prepare("INSERT INTO books set title = ?, description = ?,published = ?,image = ?,author_id = ? ")
	if err != nil {
		return nil, errors.Wrapf(err, ": rows.Prepare.Insert(%v) :", b)
	}
	defer stmt.Close()

	_, err = stmt.Exec(&b.Title, &b.Description, &b.Published, &b.Image, &b.AuthorID)
	if err != nil {
		return nil, errors.Wrapf(err, ": rows.Exec : %v", b)
	}

	data, err := json.Marshal(&b)
	if err != nil {
		return nil, errors.Wrapf(err, ": json.Marshal : %v", b)
	}

	return data, nil
}

// UpdateBook updates the specific book in the database.
func (db *DB) UpdateBook(b *Book) ([]byte, error) {

	stmt, err := db.Prepare("UPDATE books SET title = ?, description = ?, author_id = ?,image = ?, published = ? WHERE id = ?")
	if err != nil {
		return nil, errors.Wrapf(err, ": db.Prepare.Update(%v) :", b)
	}
	defer stmt.Close()

	_, err = stmt.Exec(&b.Title, &b.Description, &b.AuthorID, &b.Image, &b.Published, &b.ID)
	if err != nil {
		return nil, errors.Wrap(err, ": rows.Exec :")
	}

	data, err := json.Marshal(&b)
	if err != nil {
		return nil, errors.Wrapf(err, ": json.Marshal :")
	}

	return data, nil
}

// DeleteBook deletes the specific book from the database.
func (db *DB) DeleteBook(id string) error {

	stmt, err := db.Query("DELETE FROM books WHERE id = " + id)
	if err != nil {
		return errors.Wrapf(err, ": db.Query.Delete(%s) :", id)
	}
	defer stmt.Close()

	return nil
}
