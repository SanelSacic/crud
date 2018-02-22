// go test -v
// go test -v -run TestCreate
// go test -v -run TestUpdate
// go test -v -run TestDelete
// go test -v -run TestRetrieve
// go test -v -run TestList

package users

import (
	"testing"
	"time"

	"github.com/workspace/golang-crud/http/internal/platform/db"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCreate(t *testing.T) {

	if err := connect(); err != nil {
		t.Error(db.ErrInvalidDBProvider, err)
	}

	now := time.Now()

	insert := User{
		UserType:    1,
		FirstName:   "Mary",
		LastName:    "Jane",
		Password:    "maryjane",
		Email:       "mary.jane@gmail.com",
		Company:     "Devos",
		Image:       "jane.jpg",
		DataCreated: &now,
	}

	t.Log("Given the need to crate user")
	{
		t.Log("When using a valid user model")
		{
			user, err := Create(&insert)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a new user in the system %v", failed, insert)
			}
			t.Logf("\t%s\tShould be able to create a new user in the system : %+v", succeed, user)
		}
	}

}

func TestUpdate(t *testing.T) {

	if err := connect(); err != nil {
		t.Error(db.ErrInvalidDBProvider, err)
	}

	now := time.Now()

	update := User{
		UserID:       1,
		UserType:     2,
		FirstName:    "Max",
		LastName:     "Musterman",
		Password:     "maximito",
		Email:        "max.mus@gmail.com",
		Company:      "Devos",
		Image:        "max.jpg",
		DataModified: &now,
	}

	t.Log("Given the need to update user")
	{
		t.Log("When using a valid user model")
		{
			user, err := Update(&update)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to update user in the system %v", failed, update)
			}
			t.Logf("\t%s\tShould be able to update user in the system : %+v", succeed, user)
			t.Log(now)
		}
	}
}

func TestRetrieve(t *testing.T) {
	if err := connect(); err != nil {
		t.Error(db.ErrInvalidDBProvider, err)
	}

	userID := 1

	t.Log("Given the neet to retrieve user")
	{
		t.Log("When using a valid user model")
		{
			user, err := Retrieve(userID)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to retrieve user from the system %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to retrieve user from the system : \n %+v", succeed, user)
		}
	}
}

func TestList(t *testing.T) {
	if err := connect(); err != nil {
		t.Error(db.ErrInvalidDBProvider, err)
	}

	t.Log("Given the need to retrieve all users from system")
	{
		users, err := List()
		if err != nil {
			t.Fatalf("\t%s\tShould be able to retrieve users from system %v", failed, err)
		}
		t.Logf("\t%s\tShould be able to retrieve users from system : \n %+v", succeed, users)
	}
}

func connect() error {
	return db.NewMySql()
}
