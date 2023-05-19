package biteship

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	// Define the expected default configuration values
	expectedSecretKey := "mySecretKey"
	expectedBiteshipUrl := DefaultUrl

	// Call the DefaultConfig function with the secret key
	config := DefaultConfig(expectedSecretKey)

	// Check if the secret key is set correctly
	if config.SecretKey != expectedSecretKey {
		t.Errorf("DefaultConfig() returned unexpected secret key, got: %s, want: %s", config.SecretKey, expectedSecretKey)
	}

	// Check if the Biteship URL is set correctly
	if config.BiteshipUrl != expectedBiteshipUrl {
		t.Errorf("DefaultConfig() returned unexpected Biteship URL, got: %s, want: %s", config.BiteshipUrl, expectedBiteshipUrl)
	}
}
