package handler

import (
	"affiliate/internal/config"
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/request"
	serviceparpost "affiliate/pkg/admin/service/parPost"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// IntroduceInterface ...
type ParPostInterface interface {
	Create(c echo.Context) error
	GetList(c echo.Context) error
}

// ParPosts ...
func ParPost() ParPostInterface {
	return &parPostImpl{}
}

// introduceImpl ...
type parPostImpl struct {
}

// Create godoc
// @tags ParPost
// @summary Create
// @id parPost-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body request.ParPost true "Payload"
// @success 200 {object} nil
// @router /parPost [post]
func (p parPostImpl) Create(c echo.Context) error {
	var (
		cc   = echocustom.EchoGetCustomCtx(c)
		ctx  = cc.GetRequestCtx()
		body = cc.Get(constants.KeyPayload).(request.ParPost)
		s    = serviceparpost.ParPost()
	)
	if !config.CheckAuthen(body.Pass) {
		return cc.ResponseErr(nil, http.StatusUnauthorized, nil, "Unauthorized")

	}
	if err := s.Create(ctx, body); err != nil {
		return cc.ResponseErr(nil, http.StatusBadRequest, err, err.Error())
	}
	return cc.Response200(nil, "")
}

// Create godoc
// @tags ParPost
// @summary GetList
// @id parPost-get-list
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query request.ParPostAll true "Query"
// @success 200 {object} nil
// @router /parPost [get]
func (p parPostImpl) GetList(c echo.Context) error {
	var (
		cc  = echocustom.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
		s   = serviceparpost.ParPost()
	)
	q, ok := cc.Get(constants.KeyQuery).(request.ParPostAll)
	query := &mgquery.CommonQuery{
		Page:  q.Page,
		Limit: cc.GetLimit(q.Limit),
		SortInterface: bson.D{
			{"createdAt", -1},
		},
	}
	if !ok {
		return cc.Response400(nil, "err here")
	}

	data, err := s.GetList(ctx, query)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(data, "")
}
