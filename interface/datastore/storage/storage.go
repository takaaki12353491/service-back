package storage

import (
	"mime/multipart"
)

type Storage struct {
}

func (storage *Storage) Upload(file *multipart.FileHeader) error {
	return nil
}

func (storage *Storage) Download(file *multipart.FileHeader) error {
	return nil
}
