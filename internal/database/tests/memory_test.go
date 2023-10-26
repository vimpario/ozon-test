package tests

import (
	"ozon-test/internal/database"
	lg "ozon-test/pkg/helpers"
	"testing"
)

func TestMemoryStorageGetSecureToken(t *testing.T) {
	storage := database.NewMemoryStorage(lg.NewLogger())

	token, err := storage.getSecureToken(10)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(token) != 10 {
		t.Errorf("Expected token length to be 10, but got: %d", len(token))
	}
}
