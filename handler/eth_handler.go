package handler

import (
	"glassnode_challenge/database"

	"github.com/gofiber/fiber"

	log "github.com/sirupsen/logrus"
)

/*
 * GetETHHourlyGasFeeSpent is the api handler of the api request
 */
func GetETHHourlyGasFeeSpent(c *fiber.Ctx) {

	log.Info("API Call: GetHourlyGasFee")

	datas, err := database.DoGetHourlyGasFee()

	if err != nil {
		log.Error("dao.DoGetHourlyGasFee() ", err.Error())
		c.Status(500).JSON(fiber.Map{
			"message": "DB is not ready for connection. Please wait few more seconds " + err.Error(),
		})
		return
	}

	c.JSON(datas)
}
