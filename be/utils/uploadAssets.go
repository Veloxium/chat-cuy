package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadUserStory(ctx context.Context, filePath multipart.File) (string, error) {
	cld := database.NewCloudinary()
	res, err := cld.Upload.Upload(
		ctx,
		filePath,
		uploader.UploadParams{
			Folder: "users/story",
		})
	if err != nil {
		return "",
			fmt.Errorf("failed to upload to cloud: %w", err)
	}

	return res.SecureURL, nil
}

func UploadUserProfile(ctx context.Context, filePath multipart.File) (string, error) {
	cld := database.NewCloudinary()
	res, err := cld.Upload.Upload(
		ctx,
		filePath,
		uploader.UploadParams{
			Folder: "users/profile",
		})
	if err != nil {
		return "",
			fmt.Errorf("failed to upload to cloud: %w", err)
	}

	return res.SecureURL, nil
}
