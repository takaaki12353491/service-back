package model

import (
	"os"
	"service-back/consts"
	"service-back/validator"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Name           string `validate:"required"`
	Email          string `validate:"required,email"`
	PasswordDigest string `validate:"required"`
}

func NewUser(name string, email string, password string) (*User, error) {
	id := uuid.New().String()
	passwordDigest, _ := bcrypt.GenerateFromPassword([]byte(password+os.Getenv(consts.PASSWORD_SALT)), bcrypt.DefaultCost)
	user := &User{
		Model:          Model{ID: id},
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

func (user *User) IsAuthenticated(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password+os.Getenv(os.Getenv(consts.PASSWORD_SALT))))
	return err == nil
}
