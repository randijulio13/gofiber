package controller

import (
	"github.com/randijulio13/gofiber/internal/entity/request"
	"github.com/randijulio13/gofiber/internal/entity/response"
	"github.com/randijulio13/gofiber/internal/service"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	service service.AuthService
}

func (c *authController) Login(ctx *fiber.Ctx) error {
	request := request.LoginRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.HttpResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.HttpResponse{
			Errors: err,
			Status: false,
		})
	}

	user, err := c.service.Login(ctx, request)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.HttpResponse{
			Message: "wrong username or password",
			Status:  false,
		})
	}

	userToken, err := user.GenerateToken()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.HttpResponse{
			Message: err.Error(),
			Status:  true,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.HttpResponse{
		Message: "OK",
		Status:  true,
		Data:    userToken,
	})
}

func (c *authController) Register(ctx *fiber.Ctx) error {
	return nil
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{
		service: service,
	}
}
