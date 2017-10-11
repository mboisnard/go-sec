package sec

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRegisterShouldAddNewUserInDatabase(t *testing.T) {

	context, err := ContextInit("./authentication_test.db", "test")
	insertedID, err := RegisterUser(context, "user1", "password1")

	assert.NotNil(t, context)
	assert.Nil(t, err)
	assert.NotEqual(t, -1, insertedID)
}
