package database

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func NewCloudinary() *cloudinary.Cloudinary {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load environment" + err.Error())
	}
	uri := os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(uri)
	if err != nil {
		panic("failed to initialize cloudinary :" + err.Error())
	}
	return cld
}
