package service

import (
	"github.com/randijulio13/gofiber/internal/entity/model"
	"github.com/randijulio13/gofiber/internal/entity/request"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(ctx *fiber.Ctx, request request.LoginRequest) (*model.User, error)
}
