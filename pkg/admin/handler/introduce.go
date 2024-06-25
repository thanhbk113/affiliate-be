package handler

// import (
// 	"affiliate/internal/constants"
// 	echocustom "affiliate/internal/echo"
// 	"affiliate/internal/middleware"
// 	"affiliate/internal/util/mgquery"
// 	"affiliate/pkg/admin/model/request"
// 	servicesintroduce "affiliate/pkg/admin/service/introduce"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // IntroduceInterface ...
// type IntroduceInterface interface {
// 	Update(c echo.Context) error
// 	GetDetail(c echo.Context) error
// }

// // Introduces ...
// func Introduces() IntroduceInterface {
// 	return &introduceImpl{}
// }

// // introduceImpl ...
// type introduceImpl struct {
// }

// // Create godoc
// // @tags Introduce
// // @summary Update
// // @id introduce-update
// // @security ApiKeyAuth
// // @accept json
// // @produce json
// // @param payload body request.IntroduceBodyUpdate true "Payload"
// // @success 200 {object} nil
// // @router /introduce [put]
// func (p introduceImpl) Update(c echo.Context) error {
// 	var (
// 		cc    = echocustom.EchoGetCustomCtx(c)
// 		ctx   = cc.GetRequestCtx()
// 		body  = cc.Get(constants.KeyPayload).(request.IntroduceBodyUpdate)
// 		staff = cc.Get(constants.KeyStaff).(*middleware.User)
// 		s     = servicesintroduce.Introduce(staff)
// 	)
// 	if err := s.Update(ctx, body); err != nil {
// 		return cc.ResponseErr(nil, http.StatusBadRequest, err, err.Error())
// 	}
// 	return cc.Response200(nil, "")
// }

// // Create godoc
// // @tags Introduce
// // @summary Introduce-GetDetail
// // @id introduce-get-detail
// // @security ApiKeyAuth
// // @accept json
// // @produce json
// // @param payload query request.IntroduceAll true "Query"
// // @success 200 {object} nil
// // @router /introduce [get]
// func (p introduceImpl) GetDetail(c echo.Context) error {
// 	var (
// 		cc    = echocustom.EchoGetCustomCtx(c)
// 		ctx   = cc.GetRequestCtx()
// 		staff = &middleware.User{}
// 		s     = servicesintroduce.Introduce(staff)
// 		q     = cc.Get(constants.KeyQuery).(request.IntroduceAll)
// 		query = &mgquery.CommonQuery{
// 			ParentNameId: q.ParentNameId,
// 		}
// 	)

// 	data, err := s.GetDetail(ctx, query)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(data, "")
// }
