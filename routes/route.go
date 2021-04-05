package routes

import (
	"glassnode_challenge/handler"

	"github.com/gofiber/fiber"
)

/* Route : Routing the api request */
func Route(app *fiber.App) {

	api := app.Group("/eth")

	api.Get("/gas_hourly", handler.GetETHHourlyGasFeeSpent)
}
