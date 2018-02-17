package users

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	_, err = write.Exec(stmtInsertUser, &user.UserType, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Company, &user.Image, &user.DataCreated)
	if err != nil {
		write.Rollback()
		return nil, errors.Wrapf(err, "Insert[write.Exec] %s", stmtInsertUser)
	}

	return &user, write.Commit()
}

// Update updates a user in the database.
func Update(u *User) (*User, error) {
	write, err := db.ReadDB.Begin()
	if err != nil {
		write.Rollback()
		return nil, errors.Wrapf(err, "Update[db.ReadDB.Begin]")
	}

	now := time.Now()

	user := User{
		UserType:     u.UserType,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Password:     u.Password,
		Email:        u.Email,
		Company:      u.Company,
		Image:        u.Image,
		DataModified: &now,
	}

	_, err = write.Exec(stmtUpdateUser, &user.UserType, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Company, &user.Image, &user.UserID)
	if err != nil {
		write.Rollback()
		return nil, errors.Wrapf(err, "Update[write.Exec] %s \n %v", stmtUpdateUser, u)
	}

	return &user, write.Commit()
}
