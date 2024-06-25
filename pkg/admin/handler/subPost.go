package handler

import (
	"affiliate/internal/config"
	"affiliate/internal/constants"
	echocustom "affiliate/internal/echo"
	modelmg "affiliate/internal/model/mg"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/request"
	servicesubpost "affiliate/pkg/admin/service/subpost"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// IntroduceInterface ...
type SubPostInterface interface {
	Create(c echo.Context) error
	GetList(c echo.Context) error
}

// ParPosts ...
func SubPost() SubPostInterface {
	return &subPostImpl{}
}

// introduceImpl ...
type subPostImpl struct {
}

// Create godoc
// @tags SubPost
// @summary Create
// @id subPost-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body request.SubPost true "Payload"
// @success 200 {object} nil
// @router /subPost [post]
func (p subPostImpl) Create(c echo.Context) error {
	var (
		cc   = echocustom.EchoGetCustomCtx(c)
		ctx  = cc.GetRequestCtx()
		body = cc.Get(constants.KeyPayload).(request.SubPost)
		s    = servicesubpost.SubPost()
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
// @tags SubPost
// @summary GetList
// @id subPost-get-list
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query request.SubPostAll true "Query"
// @param id path string true "id"
// @success 200 {object} nil
// @router /subPost/{id} [get]
func (p subPostImpl) GetList(c echo.Context) error {
	var (
		cc  = echocustom.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
		s   = servicesubpost.SubPost()
		id  = cc.Get("id").(modelmg.AppID)
	)
	q, ok := cc.Get(constants.KeyQuery).(request.SubPostAll)
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

	data, err := s.GetList(ctx, query, id)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(data, "")
}
