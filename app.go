package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/fibers"
)

var app = fibers.New(NewSwagger(), fiber.Config{ErrorHandler: func(c *fiber.Ctx, err error) error {
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
