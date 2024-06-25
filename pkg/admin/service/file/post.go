package servicefile

import (
	"affiliate/internal/config"
	"affiliate/pkg/admin/model/response"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/pkg/errors"
)

func (f fileImpl) UploadImage(ctx context.Context, file multipart.File) (*response.FileResponse, error) {

	var (
		fileResponse response.FileResponse
		cfg          = config.GetENV().CloudDinary
	)

	cld, _ := cloudinary.NewFromParams(cfg.Name, cfg.ApiKey, cfg.ApiSecret)

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: cfg.FolderName,
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	fileResponse.URL = uploadResult.URL
	fileResponse.ID = uploadResult.AssetID

	defer file.Close()

	return &fileResponse, nil
}
