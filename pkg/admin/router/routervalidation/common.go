package routervalidation

import (
	echocustom "affiliate/internal/echo"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommonInterface ...
type CommonInterface interface {
	ParamID(next echo.HandlerFunc) echo.HandlerFunc
}

type commonImpl struct {
}

// ParamID ...
func (c commonImpl) ParamID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = echocustom.EchoGetCustomCtx(c)
			id = c.Param("id")
		)

		if !primitive.IsValidObjectID(id) {
			return cc.Response404(nil, "")
		}

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return cc.Response400(nil, "")
		}

		c.Set("id", objID)
		return next(c)
	}
}

func Common() CommonInterface {
	return commonImpl{}
}
