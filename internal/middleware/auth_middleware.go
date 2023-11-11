package middleware

import (
	"errors"
	"strings"

	"github.com/randijulio13/gofiber/internal/config"
	"github.com/randijulio13/gofiber/internal/entity/model"
	"github.com/randijulio13/gofiber/internal/entity/response"
	"github.com/randijulio13/gofiber/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwt(repo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(response.HttpResponse{
				Status:  false,
				Message: "unauthorized",
			})
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.HttpResponse{
				Status:  false,
				Message: "unauthorized",
			})
		}
		payload, err := validateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.HttpResponse{
				Status:  false,
				Message: "unauthorized",
			})
		}

		username := payload["username"].(string)
		user, err := repo.GetByUsername(username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.HttpResponse{
				Status:  false,
				Message: "unauthorized",
			})
		}

		c.Locals("user", user)

		return c.Next()
	}
}

func AuthUser(c *fiber.Ctx) *model.User {
	user := c.Locals("user").(*model.User)
	return user
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	config := config.GetConfig()
	secretKey := config.AccessSecret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("invalid token")
	}

	return payload, nil
}
