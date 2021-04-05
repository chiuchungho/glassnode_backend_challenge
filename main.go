package main

import (
	"glassnode_challenge/database"
	"glassnode_challenge/routes"
	"os"

	"github.com/gofiber/fiber"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
		log.Error(err.Error())
	}
	log.SetLevel(logLevel)
}

func main() {

	database.InitializeSQL()

	app := fiber.New()

	routes.Route(app)

	app.Listen(8080)
}
