package routervalidation

import (
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/pkg/admin/model/request"

	"github.com/labstack/echo/v4"
)

// ProductInterface ...
type ParPosttInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	GetList(next echo.HandlerFunc) echo.HandlerFunc
}

func ParPost() ParPosttInterface {
	return &parPostImpl{}
}

// staffImpl ...
type parPostImpl struct{}

func (p parPostImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ParPost
		)
		if err := c.Bind(&payload); err != nil {
			return cc.Response400(nil, "")
		}
		if err := payload.Validate(); err != nil {
			return cc.ValidationError(err)
		}
		c.Set(constants.KeyPayload, payload)
		return next(c)
	}
}

func (p parPostImpl) GetList(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ParPostAll
		)
		if err := c.Bind(&payload); err != nil {
			return cc.Response400(nil, "")
		}
		if err := payload.Validate(); err != nil {
			return cc.ValidationError(err)
		}
		c.Set(constants.KeyQuery, payload)
		return next(c)
	}
}
