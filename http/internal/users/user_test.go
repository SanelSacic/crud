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
			_, err := Create(&insert)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to create a new user in the system %v", failed, insert)
			}
			t.Logf("\t%s\tShould be able to create a new user in the system.", succeed)
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
			_, err := Update(&update)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to update user in the system %v", failed, update)
			}
			t.Logf("\t%s\tShould be able to update user in the system.", succeed)
		}
	}
}

func connect() error {
	return db.NewMySql("root:sonyss-a3@/made")
}
