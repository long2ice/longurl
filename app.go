package main

import (
	"embed"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/long2ice/fibers"
	"github.com/long2ice/longurl/config"
	"net/http"
)

//go:embed static/*
var static embed.FS

func initRouters(app *fibers.App) {
	app.Post("/", generateURL)
	app.Get("/:path", visitURL)
}
func initMiddlewares(app *fibers.App) {
	app.Use(
		logger.New(logger.Config{
			TimeFormat: config.ServerConfig.LogTimeFormat,
			TimeZone:   config.ServerConfig.LogTimezone,
		}),
		recover.New(),
		cors.New(),
	)

}
func CreateApp() *fibers.App {
	app := fibers.New(NewSwagger(), fiber.Config{ErrorHandler: func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if _, ok := err.(validator.ValidationErrors); ok {
			code = fiber.StatusBadRequest
		}
		err = c.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return nil
	}})
	app.Use(
		logger.New(logger.Config{
			TimeFormat: config.ServerConfig.LogTimeFormat,
			TimeZone:   config.ServerConfig.LogTimezone,
		}),
		recover.New(),
		cors.New(),
	)

	app.AfterInit(func() {
		app.Use("/", filesystem.New(filesystem.Config{
			Root:       http.FS(static),
			PathPrefix: "static",
			Browse:     true,
		}))
	})
	initRouters(app)
	initMiddlewares(app)
	return app
}
