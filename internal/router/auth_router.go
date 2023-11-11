package router

import (
	"github.com/randijulio13/gofiber/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func authRouter(router *fiber.App, authController controller.AuthController) {
	router.Post("/login", authController.Login)
}
