package sec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha512(t *testing.T) {
	hash := Sha512("Hello, World!")
	assert.Equal(t, "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387", hash)
}
