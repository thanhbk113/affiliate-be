package echocustom

import (
	"affiliate/internal/constants"
	"affiliate/internal/errorresponse"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UploadSingle ...
func UploadSingle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := EchoGetCustomCtx(c)
			// Source
			file, err := c.FormFile("file")
			if err != nil {
				return cc.ResponseErr(nil, http.StatusBadRequest, err, errorresponse.CommonKeyErrFileNotValid)
			}

			c.Set(constants.KeyFile, file)
			return next(c)
		}
	}
}
