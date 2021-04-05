package database

import (
	"os"
	"testing"
)

func TestInitializeSQL(t *testing.T) {

	os.Setenv("POSTGRES_USER", "test")
	os.Setenv("POSTGRES_PASSWORD", "test")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "eth")

	err := InitializeSQL()
	if err != nil {
		t.Errorf("Error in connecting db %v", err)
	}

}
