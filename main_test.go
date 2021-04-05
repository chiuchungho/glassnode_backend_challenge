package main

import (
	"fmt"
	"glassnode_challenge/database"
	"glassnode_challenge/routes"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
)

func TestMain(t *testing.T) {

	os.Setenv("POSTGRES_USER", "test")
	os.Setenv("POSTGRES_PASSWORD", "test")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "eth")

	database.InitializeSQL()

	app := fiber.New()

	routes.Route(app)

	resp, err := app.Test(httptest.NewRequest("GET", "/eth/gas_hourly", nil))
	fmt.Println(resp)
	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
