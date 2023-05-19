package biteship

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Prepare the test data
	testKey := "mySecretKey"

	// Call the New function to create a Biteship instance
	biteship := New(testKey)

	// Check if the Biteship instance is created
	if biteship == nil {
		t.Errorf("New() returned nil Biteship instance")
	}

	// Check if the Biteship instance has the correct Config
	config := biteship.(*Impl).Config
	if config == nil || config.SecretKey != testKey {
		t.Errorf("New() returned Biteship instance with incorrect Config")
	}
}

func TestNewWithCustomConfig(t *testing.T) {
	// Prepare the test data
	testKey := "mySecretKey"
	testURL := "https://custom.biteship.com"

	// Create a custom config
	customConfig := ConfigOption{
		SecretKey:   testKey,
		BiteshipUrl: testURL,
	}

	// Call the New function with the custom config
	biteship := New(testKey, customConfig)

	// Check if the Biteship instance is created
	if biteship == nil {
		t.Errorf("New() returned nil Biteship instance")
	}

	// Check if the Biteship instance has the correct Config
	config := biteship.(*Impl).Config
	if config == nil || config.SecretKey != testKey || config.BiteshipUrl != testURL {
		t.Errorf("New() returned Biteship instance with incorrect Config")
	}
}
