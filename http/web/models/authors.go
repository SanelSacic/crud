package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Description string `json:"description"`
	Birth       string `json:"birth"`
}

// ListAuthors retrieves a list of existing authors in the database.
func (db *DB) ListAuthors() ([]byte, error) {
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, errors.Wrap(err, ": conn.Query :")
	}
	defer rows.Close()

	var authors []*Author

	for rows.Next() {
		var a Author
		err := rows.Scan(&a.ID, &a.Name, &a.LastName, &a.Description, &a.Birth)
		if err != nil {
			return nil, errors.Wrapf(err, ": rows.Next: %v", a)
		}
		authors = append(authors, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, ": rows.Err :")
	}

	var data []byte

	if data, err = json.Marshal(&authors); err != nil {
		return nil, errors.Wrap(err, ": json.Marhal :")
	}

	return data, nil
}

// RetrieveAuthor gets the specific book from the database.
func (db *DB) RetrieveAuthor(id string) ([]byte, error) {

	rows, err := db.Query("SELECT * FROM authors WHERE id = " + id)
	if err != nil {
		return nil, errors.Wrapf(err, ": conn.Query.Retrieve(%s)", rows)
	}
	defer rows.Close()

	var res []*Author

	if rows.Next() {
		var a Author
		err := rows.Scan(&a.ID, &a.Name, &a.LastName, &a.Description, &a.Birth)
		if err != nil {
			return nil, errors.Wrapf(err, ": rows.Scann : %v", a)
		}
		res = append(res, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, ": rows.Err :")
	}

	var data []byte

	if data, err = json.Marshal(&res); err != nil {
		return nil, errors.Wrapf(err, ": json.Marshal : %d", data)
	}

	return data, nil
}

// CreateAuthor creates a new Author in the database.
func (db *DB) CreateAuthor(a *Author) ([]byte, error) {

	stmt, err := db.Prepare("INSERT INTO authors SET name = ?, lastname = ?, description = ?,birth = ?")
	if err != nil {
		return nil, errors.Wrapf(err, ": rows.Prepare.Insert(%s) :", a)
	}
	defer stmt.Close()

	_, err = stmt.Exec(&a.Name, &a.LastName, &a.Description, &a.Birth)
	if err != nil {
		return nil, errors.Wrapf(err, ": rows.Exec : %v", a)
	}

	data, err := json.Marshal(&a)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Marshal : %v", a)
	}

	return data, nil
}

// UpdateAuthor updates the specific author in the database.
func (db *DB) UpdateAuthor(a *Author) ([]byte, error) {

	stmt, err := db.Prepare("UPDATE authors SET name = ?, lastname = ?, description = ?,birth = ? WHERE id = ?")
	if err != nil {
		return nil, errors.Wrapf(err, ": db.Prepare.Update(%+v) :", a)
	}
	defer stmt.Close()
	_, err = stmt.Exec(&a.Name, &a.LastName, &a.Description, &a.Birth, &a.ID)
	if err != nil {
		return nil, errors.Wrapf(err, ": rows.Exec : %v", a)
	}

	data, err := json.Marshal(&a)
	if err != nil {
		return nil, errors.Wrap(err, "json.Marshal")
	}

	return data, nil
}

// DeleteAuthor deletes the specific author from the database.
func (db *DB) DeleteAuthor(id string) error {

	stmt, err := db.Query("DELETE FROM authors WHERE id = " + id)
	if err != nil {
		return errors.Wrapf(err, ": db.Query.Delete(%s) :", id)
	}
	defer stmt.Close()

	return nil
}

// IsValid validates author.
func (a *Author) IsValid() bool {
	switch {
	case len(a.Name) == 0:
		return false

	case len(a.LastName) == 0:
		return false

	case len(a.Description) == 0:
		return false

	case len(a.Birth) == 0:
		return false
	}

	return true
}
