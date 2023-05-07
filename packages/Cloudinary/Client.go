package cloudinary

import (
	"context"
	"fmt"
	gongaCloudinary "gonga/contracts/Cloudinary"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"gonga/config"
)

type CloudinaryClient struct {
	Config *config.CloudinaryConfig
	Client *cloudinary.Cloudinary
	Ctx    context.Context
}

func NewCloudinaryClient() *CloudinaryClient {
	config := config.LoadCloudinaryConfig()
	// Add your Cloudinary product environment credentials.

	cld, _ := cloudinary.NewFromParams(config.CloudName, config.APIKey, config.APISecret)
	cld.Config.URL.Secure = true
	ctx := context.Background()

	return &CloudinaryClient{Config: config, Client: cld, Ctx: ctx}
}

func (c *CloudinaryClient) GetConfig() *config.CloudinaryConfig {
	return c.Config
}

func (c *CloudinaryClient) UploadImage(file multipart.File, publicID string) (*gongaCloudinary.CloudinaryResponse, error) {
	// Set upload options
	uploadOptions := uploader.UploadParams{
		PublicID:       publicID,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	}

	// Upload image
	uploadResult, err := c.Client.Upload.Upload(c.Ctx, file, uploadOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %v", err)
	}

	// Return response
	return &gongaCloudinary.CloudinaryResponse{
		URL:      uploadResult.SecureURL,
		Type:     uploadResult.Format,
		Filename: uploadResult.OriginalFilename,
		Size:     uploadResult.Bytes,
	}, nil
}

func (c *CloudinaryClient) UploadVideo(filepath string, publicID string) (*gongaCloudinary.CloudinaryResponse, error) {
	// Set upload options
	uploadOptions := uploader.UploadParams{
		PublicID:     publicID,
		Overwrite:    api.Bool(true),
		ResourceType: "video",
	}

	// Upload video
	uploadResult, err := c.Client.Upload.Upload(c.Ctx, filepath, uploadOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to upload video: %v", err)
	}

	// Return response
	return &gongaCloudinary.CloudinaryResponse{
		URL:      uploadResult.SecureURL,
		Type:     uploadResult.Format,
		Filename: uploadResult.OriginalFilename,
		Size:     uploadResult.Bytes,
	}, nil
}

func (c *CloudinaryClient) DeleteFile(publicID string) (*gongaCloudinary.CloudinaryResponse, error) {
	// Set delete options
	deleteOptions := uploader.DestroyParams{
		PublicID: publicID,
	}

	// Delete file
	_, err := c.Client.Upload.Destroy(c.Ctx, deleteOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to delete file: %v", err)
	}

	// Return response
	return &gongaCloudinary.CloudinaryResponse{
		URL:     "",
		Type:    "",
		Message: "File deleted succesfully.",
	}, nil
}
