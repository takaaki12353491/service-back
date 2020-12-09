package storage

import (
	"io"
	"mime/multipart"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	dir = "./image/"
)

func Upload(fh *multipart.FileHeader) error {
	if fh == nil {
		log.Info("File doesn't exist")
		return nil
	}
	img, err := fh.Open()
	if err != nil {
		log.Error(err)
		return err
	}
	defer img.Close()
	dst, err := os.Create(dir + fh.Filename)
	if err != nil {
		log.Error(err)
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, img)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func Download(file *multipart.FileHeader) error {
	return nil
}
