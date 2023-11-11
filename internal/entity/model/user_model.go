package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"ID" gorm:"not null;primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email,omitempty" gorm:"default:null"`
	gorm.Model
}

type UserToken struct {
	UserID       string `json:"-"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (u *User) GenerateToken() (*UserToken, error) {
	claims := jwt.MapClaims{
		"username": u.Username,
		"name":     u.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	userToken := UserToken{
		UserID:      u.ID,
		AccessToken: tokenString,
	}

	return &userToken, nil
}

func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
