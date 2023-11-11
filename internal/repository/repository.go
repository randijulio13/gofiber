package repository

import "github.com/randijulio13/gofiber/internal/entity/model"

type UserRepository interface {
	GetAllUser()
	StoreUser(*model.User) error
	GetByUsername(string) (*model.User, error)
}
