package service

import (
	"errors"

	"github.com/randijulio13/gofiber/internal/entity/model"
	"github.com/randijulio13/gofiber/internal/entity/request"
	"github.com/randijulio13/gofiber/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type authService struct {
	repo repository.UserRepository
}

func (s *authService) Login(ctx *fiber.Ctx, request request.LoginRequest) (*model.User, error) {
	user, err := s.repo.GetByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if err := user.CheckPassword(request.Password); err != nil {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}
