package routervalidation

import (
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/pkg/admin/model/request"

	"github.com/labstack/echo/v4"
)

// ProductInterface ...
type SubPosttInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	GetList(next echo.HandlerFunc) echo.HandlerFunc
	Update(next echo.HandlerFunc) echo.HandlerFunc
}

func SubPost() SubPosttInterface {
	return &subPostImpl{}
}

// subPostImpl ...
type subPostImpl struct{}

func (p subPostImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.SubPost
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

func (p subPostImpl) GetList(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.SubPostAll
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

func (p subPostImpl) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.SubPostUpdate
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
