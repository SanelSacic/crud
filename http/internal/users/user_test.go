package users

import (
	"testing"

	"github.com/workspace/golang-crud/http/internal/platform/db"
)

const suceed = "\u2713"
const failed = "\u2717"

func TestCreate(t *testing.T) {

	if err := connect(); err != nil {
		t.Error(ErrInvalidDBProvider, err)
	}
}

func connect() error {
	return db.NewMySQL("root:sonyss-a3@/made")
}
