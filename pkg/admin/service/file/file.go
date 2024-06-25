package servicefile

import (
	"affiliate/internal/middleware"
	"affiliate/pkg/admin/model/response"
	"context"
	"mime/multipart"
)

type FileInterface interface {
	//Upload image file to cloudinary
	UploadImage(ctx context.Context, file multipart.File) (*response.FileResponse, error)
}

type fileImpl struct {
	User *middleware.User
}

// Package ...
func File(user *middleware.User) FileInterface {
	return &fileImpl{
		User: user,
	}
}
