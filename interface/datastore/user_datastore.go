package datastore

import (
	"service-back/domain/model"
	"service-back/interface/datastore/database"
)

type UserDatastore struct {
	database *database.UserDatabase
}

func NewUserDatastore() *UserDatastore {
	return &UserDatastore{
		database: database.NewUserDatabase(),
	}
}

func (ds *UserDatastore) FindByNameOrEmail(name, email string) (*model.User, error) {
	return ds.database.FindByNameOrEmail(name, email)
}

func (ds *UserDatastore) FindByID(id string) (*model.User, error) {
	return ds.database.FindByID(id)
}

func (ds *UserDatastore) Store(user *model.User) error {
	return ds.database.Store(user)
}

func (ds *UserDatastore) Update(user *model.User) error {
	return ds.database.Update(user)
}
