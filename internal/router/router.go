package router

import (
	"github.com/randijulio13/gofiber/internal/controller"
	"github.com/randijulio13/gofiber/internal/entity/response"
	"github.com/randijulio13/gofiber/internal/middleware"
	"github.com/randijulio13/gofiber/internal/repository"
	"github.com/randijulio13/gofiber/internal/service"
	"github.com/randijulio13/gofiber/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func Route(router *fiber.App) {
	db := database.GetDatabase()
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	authRouter(router, authController)

	router.Use(middleware.VerifyJwt(userRepository))
	router.Get("/", func(c *fiber.Ctx) error {
		user := middleware.AuthUser(c)
		return c.Status(fiber.StatusOK).JSON(response.HttpResponse{
			Status: true,
			Data:   user,
		})
	})
}
