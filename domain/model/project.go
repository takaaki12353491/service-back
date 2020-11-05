package model

import (
	"mime/multipart"
)

type Project struct {
	Model
	CommunityID string
	Community   Community
	OwnerID     string
	Owner       User
	Name        string
	Description string
	LogoURL     string
	Logo        *multipart.FileHeader
	HeaderURL   string
	Header      *multipart.FileHeader
	Members     []User
}
