package database

import (
	"service-back/domain/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type CommunityDatabase struct {
	*gorm.DB
}

func NewCommunityDatabase() *CommunityDatabase {
	return &CommunityDatabase{NewConnection()}
}

func (db *CommunityDatabase) FindAll() ([]model.Community, error) {
	communities := []model.Community{}
	tx := db.Set("gorm:auto_preload", true).Begin()
	err := tx.Find(&communities).Commit().Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return communities, nil
}

func (db *CommunityDatabase) Store(community *model.Community) error {
	err := db.Create(community).Error
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
