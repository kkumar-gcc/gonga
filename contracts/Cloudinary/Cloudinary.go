package cloudinary

import (
	"gonga/config"
	"mime/multipart"
)

type CloudinaryService interface {
	GetConfig() *config.CloudinaryConfig
	UploadImage(file multipart.File, publicID string) (*CloudinaryResponse, error)
	UploadVideo(filepath string, publicID string) (*CloudinaryResponse, error)
	DeleteFile(publicID string) (*CloudinaryResponse, error)
}

type CloudinaryResponse struct {
	URL      string `json:"url"`
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Size     int  `json:"size"`
	Message  string `json:"message"`
}
