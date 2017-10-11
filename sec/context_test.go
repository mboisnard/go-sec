package sec

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestContextInit(t *testing.T) {
	salt := "test"
	context, err := ContextInit("./test.db", nil, salt)

	assert.Nil(t, err)
	assert.NotNil(t, context)
	assert.Equal(t, salt, context.salt)
	assert.NotNil(t, context.connection)
}
