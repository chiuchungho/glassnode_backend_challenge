package database

import (
	"os"
	"testing"
)

func TestDoGetHourlyGasFee(t *testing.T) {

	os.Setenv("POSTGRES_USER", "test")
	os.Setenv("POSTGRES_PASSWORD", "test")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "eth")

	err := InitializeSQL()
	if err != nil {
		t.Errorf("Error in connecting db %v", err)
	}

	datas, err := DoGetHourlyGasFee()
	if err != nil {
		t.Errorf("Error in DoGetHourlyGasFee %v", err)
	}

	if len(datas) != 24 {
		t.Errorf("Error in DoGetHourlyGasFee result. Slice length != 24")
	}

}
