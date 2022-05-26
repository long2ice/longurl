package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/long2ice/longurl/config"
)

func init() {
	app.Use(
		logger.New(logger.Config{
			TimeFormat: config.ServerConfig.LogTimeFormat,
			TimeZone:   config.ServerConfig.LogTimezone,
		}),
		recover.New(),
		cors.New(),
	)
}
