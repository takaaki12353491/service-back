package database

import (
	"service-back/domain/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type UserDatabase struct {
	*gorm.DB
}

func NewUserDatabase() *UserDatabase {
	return &UserDatabase{NewConnection()}
}

func (db *UserDatabase) FindByNameOrEmail(name, email string) (*model.User, error) {
	user := new(model.User)
	err := db.Where("name = ? OR email = ?", name, email).First(user).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func (db *UserDatabase) FindByID(id string) (*model.User, error) {
	user := new(model.User)
	user.ID = id
	err := db.First(user).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func (db *UserDatabase) Store(user *model.User) error {
	err := db.Create(user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (db *UserDatabase) Update(user *model.User) error {
	err := db.Save(user).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
