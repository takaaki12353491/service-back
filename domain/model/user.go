package model

import (
	"service-back/validator"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name           string `validate:"required"`
	Email          string `validate:"required,email"`
	PasswordDigest string `validate:"required"`
}

func NewUser(name string, email string, password string) (*User, error) {
	id := uint(uuid.New().ID())
	passwordDigest, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &User{
		Model:          gorm.Model{ID: id},
		Name:           name,
		Email:          email,
		PasswordDigest: string(passwordDigest),
	}
	err := validator.Validate(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
