package sec

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestContextInit(t *testing.T) {
	context, err := ContextInit("./test.db")

	assert.Nil(t, err)
	assert.NotNil(t, context)
	assert.NotNil(t, context.connection)
}
