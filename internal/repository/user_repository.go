package repository

import (
	"errors"

	"github.com/randijulio13/gofiber/internal/entity/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) StoreUser(user *model.User) error {
	user.ID = uuid.NewString()
	err := repo.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetAllUser() {

}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	q := r.db.Where("username = ?", username).Find(&user)

	if q.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
