package inputdata

import (
	"mime/multipart"
)

type Community struct {
	UserID      string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Header      *multipart.FileHeader
	Icon        *multipart.FileHeader
}

type UpdateCommunity struct {
	ID string
	Community
}
