package sec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewSessionInDatabase(t *testing.T) {

	context, err := ContextInit("./sessions_test.db", nil, "")
	assert.Nil(t, err)

	insertedID, err := RegisterUser(context, "user1", "password1")
	assert.Nil(t, err)
	assert.NotEqual(t, -1, insertedID)

	token, err := CreateNewSession(context, insertedID)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
