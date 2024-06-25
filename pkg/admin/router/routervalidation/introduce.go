package routervalidation

import (
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/pkg/admin/model/request"

	"github.com/labstack/echo/v4"
)

// IntroduceInterface ...
type IntroduceInterface interface {
	Update(next echo.HandlerFunc) echo.HandlerFunc
	All(next echo.HandlerFunc) echo.HandlerFunc
}

func Introduces() IntroduceInterface {
	return &introduceImpl{}
}

// introduceImpl ...
type introduceImpl struct{}

func (p introduceImpl) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.IntroduceBodyUpdate
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

func (p introduceImpl) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc      = echocustom.EchoGetCustomCtx(c)
			payload request.IntroduceAll
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
