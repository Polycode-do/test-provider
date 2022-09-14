package client

import (
	"os"
	"testing"
)

func TestNewClientWithAuth(t *testing.T) {
	username := "admin@gmail.com"
	password := "12345678"
	if os.Getenv("POLYCODE_USERNAME") != "" {
		username = os.Getenv("POLYCODE_USERNAME")
	}
	if os.Getenv("POLYCODE_PASSWORD") != "" {
		password = os.Getenv("POLYCODE_PASSWORD")
	}

	_, err := NewClient(nil, &username, &password)
	if err != nil {
		t.Errorf("Error creating client: %s", err)
	}
}

func TestNewClientWithoutAuth(t *testing.T) {
	_, err := NewClient(nil, nil, nil)
	if err != nil {
		t.Errorf("Error creating client: %s", err)
	}
}
