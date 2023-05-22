package biteship

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	testKey := "mySecretKey"

	biteship := New(
		WithSecret(testKey),
	)

	assert.NotNil(t, biteship)
	config := biteship.(*Client)
	assert.NotNil(t, config)
	assert.Equal(t, config.SecretKey, testKey)
}

func TestNewWithCustomConfig(t *testing.T) {
	testKey := "mySecretKey"
	testURL := "https://custom.biteship.com"

	biteship := New(WithSecret(testKey), WithUrl(testURL))

	assert.NotNil(t, biteship)
	config := biteship.(*Client)
	assert.NotNil(t, config)
	assert.Equal(t, config.SecretKey, testKey)
	assert.Equal(t, config.BiteshipUrl, testURL)
}
