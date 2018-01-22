package sec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha512(t *testing.T) {
	hash := Sha512("Hello, World!")
	assert.Equal(t, "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387", hash)
}

func TestGenerateKeypair(t *testing.T) {
	key, err := GenerateKeypair(4096)

	assert.NotNil(t, key)
	assert.Nil(t, err)

	assert.Equal(t, key.E, key.PublicKey.E)
	assert.Equal(t, key.N, key.PublicKey.N)
}

func TestGenerateKeypairShouldThrowIfLessThan4096Bits(t *testing.T) {
	key, err := GenerateKeypair(2048)

	assert.Nil(t, key)
	assert.NotNil(t, err)
}

func TestEncodePrivateKey(t *testing.T) {
	key, _ := GenerateKeypair(4096)
	pem, err := EncodePrivateKey(key)

	assert.NotNil(t, pem)
	assert.Nil(t, err)
}

func TestEncodePrivateKeyShouldThrowWhenKeyIsNil(t *testing.T) {
	pem, err := EncodePrivateKey(nil)

	assert.Nil(t, pem)
	assert.NotNil(t, err)
}

func TestEncodePublicKey(t *testing.T) {
	key, _ := GenerateKeypair(4096)
	pem, err := EncodePublicKey(&key.PublicKey)

	assert.NotNil(t, pem)
	assert.Nil(t, err)
}

func TestEncodePublicKeyShouldThrowWhenKeyIsNil(t *testing.T) {
	pem, err := EncodePublicKey(nil)

	assert.Nil(t, pem)
	assert.NotNil(t, err)
}
