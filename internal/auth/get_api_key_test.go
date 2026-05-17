package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "abc123" {
		t.Fatalf("expected 'abc123', got %s", key)
	}
}

func TestGetAPIKey_InvalidHeader(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "Bearer abc123")

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}