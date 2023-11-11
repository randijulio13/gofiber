package controller

import "github.com/gofiber/fiber/v2"

type ControllerFunc func(ctx *fiber.Ctx) error

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}
