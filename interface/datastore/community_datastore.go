package datastore

import (
	"service-back/domain/model"
	"service-back/interface/datastore/database"
)

type CommunityDatastore struct {
	database *database.CommunityDatabase
}

func NewCommunityDatastore() *CommunityDatastore {
	return &CommunityDatastore{
		database: database.NewCommunityDatabase(),
	}
}

func (ds *CommunityDatastore) FindAll() ([]model.Community, error) {
	return ds.database.FindAll()
}

func (ds *CommunityDatastore) FindByID(id string) (*model.Community, error) {
	return ds.database.FindByID(id)
}

func (ds *CommunityDatastore) Store(community *model.Community) error {
	return ds.database.Store(community)
}
