package handler

import (
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/internal/errorresponse"
	"affiliate/internal/middleware"
	"affiliate/internal/util"
	servicefile "affiliate/pkg/admin/service/file"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

// FileInterface ...
type FileInterface interface {
	// UploadImage ...
	UploadImage(c echo.Context) error
}

// File ...
func File() FileInterface {
	return &fileImpl{}
}

// fileImpl ...
type fileImpl struct {
}

// Create godoc
// @tags Files
// @summary Upload image
// @id file-upload-image
// @security ApiKeyAuth
// @accept json
// @produce json
// @param file formData file true "File"
// @success 200 {object} nil
// @router /file/photo [post]
func (p fileImpl) UploadImage(c echo.Context) error {
	var (
		cc  = echocustom.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
	)
	f, ok := cc.Get(constants.KeyFile).(*multipart.FileHeader)

	if !ok {
		return cc.ResponseErr(nil, http.StatusBadRequest, nil, errorresponse.CommonKeyErrFileNotValid)
	}
	staff := &middleware.User{}

	s := servicefile.File(staff)

	//create file
	filePath, _ := util.CreateFile(f)

	//resize image
	err := util.ResizeAndSaveImage(filePath, constants.ImageWidth, constants.ImageHeight)

	if err != nil {
		return cc.ResponseErr(nil, http.StatusInternalServerError, err, errorresponse.CommonKeyErrResizeImage)
	}

	//convert to multipart.File
	mtf, err := util.ConvertFileHeaderToMultipartFile(filePath)

	if err != nil {
		return cc.ResponseErr(nil, http.StatusInternalServerError, err, errorresponse.CommonKeyErrReadFile)
	}

	data, err := s.UploadImage(ctx, *mtf)
	if err != nil {
		return cc.ResponseErr(nil, http.StatusInternalServerError, err, errorresponse.CommonKeyErrUploadFile)
	}
	//delete file
	_ = util.DeleteFile(filePath)

	return cc.Response200(data, "")
}
