package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"long2ice/longurl/config"
)

var app = fiber.New()

func main() {
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.ServerConfig.Host, config.ServerConfig.Port)))
}
