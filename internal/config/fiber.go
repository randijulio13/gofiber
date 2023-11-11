package config

import (
	"github.com/gofiber/fiber/v2"
)

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       config.AppName,
	}
}
