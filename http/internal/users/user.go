package users

import (
	"time"

	"github.com/pkg/errors"
	"github.com/workspace/golang-crud/http/internal/platform/db"
)

const (
	stmtInsertUser   = "INSERT INTO users (userType ,firstName, lastName, password, email, company,image, dataCreated) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmtUpdateUser   = "UPDATE users set userType = ?, firstName = ?, lastName = ?, password = ?, email = ?, company = ?, image = ? WHERE id = ?"
	stmtDeleteUser   = "DELETE FROM users WHERE id = ?"
	stmtRetrieveUser = "SELECT id,userType,firstName,lastName,email,company,image FROM users WHERE id = ?"
	stmtListUsers    = "SELECT id,userType,firstName,lastName,email,company,image FROM users"
)

// Create inserts a new user into the database.
func Create(u *User) (*User, error) {
	write, err := db.ReadDB.Begin()
	if err != nil {
		write.Rollback()
		return nil, errors.Wrap(err, "Insert[db.ReadDB.Begin]")
	}

	now := time.Now()

	user := User{
		UserType:    u.UserType,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Password:    u.Password,
		Email:       u.Email,
		Company:     u.Company,
		Image:       u.Image,
		DataCreated: &now,
	}

	_, err = write.Exec(stmtInsertUser, &user.UserType, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Company, &user.Image, &user.DateCreated)
	if err != nil {
		write.Rollback()
		return nil, errors.Wrapf(err, "Insert[write.Exec] %s", stmtInsertUser)
	}
}
