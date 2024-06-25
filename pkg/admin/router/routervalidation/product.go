package routervalidation

import (
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/pkg/admin/model/request"

	"github.com/labstack/echo/v4"
)

// ProductInterface ...
type ProductInterface interface {
	Create(next echo.HandlerFunc) echo.HandlerFunc
	All(next echo.HandlerFunc) echo.HandlerFunc
	AllUser(next echo.HandlerFunc) echo.HandlerFunc
	Update(next echo.HandlerFunc) echo.HandlerFunc
}

func Product() ProductInterface {
	return &productImpl{}
}

// staffImpl ...
type productImpl struct{}

func (p productImpl) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ProductBody
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

func (p productImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ProductAll
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

func (p productImpl) AllUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ProductAllUser
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

func (p productImpl) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.ProductBodyUpdate
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
