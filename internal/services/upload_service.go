package services

import (
	"appointment-api/internal/config"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadService interface {
	UploadImage(file multipart.File, filename string, folder string) (*UploadResult, error)
	DeleteImage(publicID string) error
}

type UploadResult struct {
	PublicID  string `json:"public_id"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Format    string `json:"format"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Bytes     int    `json:"bytes"`
}

type uploadService struct {
	cloudinary *cloudinary.Cloudinary
}

func NewUploadService(cfg *config.Config) (UploadService, error) {
	// 2.x: Bağlantı için CLOUDINARY_URL kullan
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}
	return &uploadService{
		cloudinary: cld,
	}, nil
}

func (s *uploadService) UploadImage(file multipart.File, filename string, folder string) (*UploadResult, error) {
	ctx := context.Background()

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	if !allowedExts[ext] {
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	// Generate unique public ID
	publicID := fmt.Sprintf("%s/%s_%d", folder,
		strings.TrimSuffix(filename, filepath.Ext(filename)),
		time.Now().Unix())

	uploadParams := uploader.UploadParams{
		PublicID:     publicID,
		Folder:       folder,
		ResourceType: "image",
		Overwrite:    api.Bool(true),
	}

	uploadResult, err := s.cloudinary.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Cloudinary: %v", err)
	}

	result := &UploadResult{
		PublicID:  uploadResult.PublicID,
		URL:       uploadResult.URL,
		SecureURL: uploadResult.SecureURL,
		Format:    uploadResult.Format,
		Width:     uploadResult.Width,
		Height:    uploadResult.Height,
		Bytes:     uploadResult.Bytes,
	}

	return result, nil
}

func (s *uploadService) DeleteImage(publicID string) error {
	ctx := context.Background()
	if publicID == "" {
		return fmt.Errorf("public ID is required")
	}
	deleteParams := uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	}
	_, err := s.cloudinary.Upload.Destroy(ctx, deleteParams)
	if err != nil {
		return fmt.Errorf("failed to delete from Cloudinary: %v", err)
	}
	return nil
}
