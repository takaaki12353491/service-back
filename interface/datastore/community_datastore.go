package datastore

import (
	"service-back/domain/model"
	"service-back/interface/datastore/database"
	"service-back/interface/datastore/storage"

	log "github.com/sirupsen/logrus"
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
	err := ds.database.Store(community)
	if err != nil {
		log.Error(err)
		return err
	}
	storage.Upload(community.Header)
	storage.Upload(community.Icon)
	return nil
}
