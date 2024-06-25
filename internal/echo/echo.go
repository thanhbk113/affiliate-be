package echocustom

import (
	"affiliate/internal/constants"
	"affiliate/internal/locale"
	"affiliate/internal/module/logger"
	"affiliate/internal/util/pmg"
	"affiliate/internal/util/ptime"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EchoGetCustomCtx ...
func EchoGetCustomCtx(c echo.Context) *EchoCustomCtx {
	return &EchoCustomCtx{c}
}

// EchoCustomCtx custom echo context
type EchoCustomCtx struct {
	echo.Context
}

// GetHeaderKey ...
func (c *EchoCustomCtx) GetHeaderKey(k string) string {
	return c.Request().Header.Get(k)
}

// GetString ...
func (c *EchoCustomCtx) GetString(key string) string {
	v := c.Get(key)
	res, _ := v.(string)
	return res
}

// GetVersionCode ...
func (c *EchoCustomCtx) GetVersionCode() int64 {
	versionCode, _ := strconv.Atoi(c.QueryParam("versionCodeFrom"))
	return int64(versionCode)
}

// GetPageQuery ...
func (c *EchoCustomCtx) GetPageQuery() int64 {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	return int64(page)
}

// GetLimitQuery ...
func (c *EchoCustomCtx) GetLimitQuery() int64 {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		if limit == -1 {
			limit = 0
		} else {
			limit = 20
		}
	}

	return int64(limit)
}

// GetIntQuery ...
func (c *EchoCustomCtx) GetIntQuery(key string) int {
	n, _ := strconv.Atoi(c.QueryParam("limit"))
	return n
}

// GetLang ...
func (c *EchoCustomCtx) GetLang() string {
	lang := c.GetHeaderKey(constants.HeaderAcceptLanguage)
	if lang != constants.LangVI {
		lang = constants.LangEN
	}
	return lang
}

// GetCurrentUserID ...
func (c *EchoCustomCtx) GetCurrentUserID() (id primitive.ObjectID) {
	token := c.Get("user")
	if token == nil {
		return
	}

	data, ok := token.(*jwt.Token)
	if !ok {
		return
	}

	m, ok := data.Claims.(jwt.MapClaims)
	if ok && data.Valid {
		s, ok := m["_id"].(string)
		if ok && s != "" {
			id = pmg.GetAppIDFromHex(s)
		}
	}

	return id
}

// GetToken ...
func (c *EchoCustomCtx) GetToken() string {
	token := strings.Split(c.GetHeaderKey(constants.HeaderAuthorization), " ")
	if len(token) == 1 {
		return ""
	}
	return token[1]
}

// GetRequestCtx get request context
func (c *EchoCustomCtx) GetRequestCtx() context.Context {
	return c.Request().Context()
}

// GetListQueryParam ...
func (c *EchoCustomCtx) GetListQueryParam(filed string) (res []string) {
	data := c.QueryParam(filed)
	if data == "" {
		return
	}
	res = strings.Split(data, ",")
	return
}

// GetUserPlatform ...
func (c *EchoCustomCtx) GetUserPlatform() string {
	return strings.ToLower(c.Request().Header.Get(constants.HeaderOSName))
}

// GetOsSdkVersion ...
func (c *EchoCustomCtx) GetOsSdkVersion() string {
	return strings.ToLower(c.Request().Header.Get("OS-SDK"))
}

// GetAppVersion ...
func (c *EchoCustomCtx) GetAppVersion() string {
	return strings.ToLower(c.Request().Header.Get("App-Version"))
}

// GetOsVersion ...
func (c *EchoCustomCtx) GetOsVersion() string {
	return strings.ToLower(c.Request().Header.Get("OS-Version"))
}

// GetOsName ...
func (c *EchoCustomCtx) GetOsName() string {
	return strings.ToLower(c.Request().Header.Get("OS-Name"))
}

// GetDeviceId ...
func (c *EchoCustomCtx) GetDeviceId() string {
	return strings.ToLower(c.Request().Header.Get("DeviceId"))
}

// GetAppIDFromQuery ...
func (c *EchoCustomCtx) GetAppIDFromQuery(key string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(c.QueryParam(key))
	return id
}

// GetVersion ...
func (c *EchoCustomCtx) GetVersion() string {
	return strings.ToLower(c.Request().Header.Get("version"))
}

// GetRequestURI ...
func (c *EchoCustomCtx) GetRequestURI() string {
	return c.Request().RequestURI
}

// GetBodyRaw ...
func (c *EchoCustomCtx) GetBodyRaw() []byte {
	b, _ := ioutil.ReadAll(c.Request().Body)
	return b
}

// Response200 response success
func (c *EchoCustomCtx) Response200(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default200
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusOK)
	sendResponse(c, resp)
	return nil
}

// Response400 bad request
func (c *EchoCustomCtx) Response400(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default400
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusBadRequest)
	sendResponse(c, resp)
	return nil
}

// Response401 Unauthorized
func (c *EchoCustomCtx) Response401(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default401
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusUnauthorized)
	sendResponse(c, resp)
	return nil
}

// ValidationError ...
func (c *EchoCustomCtx) ValidationError(err error) error {
	logger.Info("ValidationError", logger.LogData{
		Source:  "",
		Message: err.Error(),
		Data:    nil,
	})
	cc := EchoGetCustomCtx(c)
	return cc.Response400(nil, getMessage(err))
}

func getMessage(err error) string {
	err1, ok := err.(validation.Errors)
	if !ok {
		err2, ok := err.(validation.ErrorObject)
		if ok {
			return err2.Message()
		}
		return err.Error()
	}
	for _, item := range err1 {
		if item == nil {
			continue
		}
		return getMessage(item)
	}
	return err.Error()
}

// Response403 Forbidden
func (c *EchoCustomCtx) Response403(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default403
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusForbidden)
	sendResponse(c, resp)
	return nil
}

// Response400 bad request
func (c *EchoCustomCtx) ResponseErr(data interface{}, statusCode int, err error, message string) error {

	logger.GetZapLogger().Error(fmt.Sprintf("ResponseErr: %+v \n", errors.WithStack(err)))

	c.JSON(statusCode, echo.Map{
		"data":    data,
		"message": message,
		"code":    statusCode,
	})

	return nil
}

// GetLimit ...
func (c *EchoCustomCtx) GetLimit(param int64) int64 {
	if param <= 0 {
		return 20
	}
	return param
}

// Response404 not found
func (c *EchoCustomCtx) Response404(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default404
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusNotFound)
	sendResponse(c, resp)
	return nil
}

// DefineResponse ...
func (c *EchoCustomCtx) DefineResponse(code int, data interface{}, msgKey string) error {
	resp := getResponse(c.GetLang(), data, msgKey, code)
	sendResponse(c, resp)
	return nil
}

// ActionType ...
type ActionType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Text  string `json:"text"`
}

// Response ...
type Response struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
}

func sendResponse(c echo.Context, data Response) {
	if err := c.JSON(data.HTTPCode, echo.Map{
		"data":    data.Data,
		"message": data.Message,
		"code":    data.Code,
	}); err != nil {
		logger.Error("echo.sendResponse", logger.LogData{
			Source:  "",
			Message: err.Error(),
			Data:    data,
		})
	}
}

func getResponse(lang string, data interface{}, messageKey string, httpCode int) Response {
	if data == nil {
		data = echo.Map{}
	}

	var respInfo locale.Locale
	respInfo = locale.GetByKey(lang, messageKey)
	respInfo.Message.GetDisplay(lang)
	if respInfo.Key == "NotFound" && messageKey != "" {
		respInfo.Message.Display = messageKey
	}
	return Response{
		HTTPCode: httpCode,
		Data:     data,
		Message:  respInfo.Message.Display,
		Code:     respInfo.Code,
	}
}

// ConvertQueryToInt64 ...
func (c *EchoCustomCtx) ConvertQueryToInt64(queryParam string) int64 {
	number, _ := strconv.Atoi(c.QueryParam(queryParam))
	return int64(number)
}

// ConvertQueryToFloat64 ...
func (c *EchoCustomCtx) ConvertQueryToFloat64(queryParam string) float64 {
	number, _ := strconv.ParseFloat(c.QueryParam(queryParam), 64)
	return number
}

// GetPlatform ...
func (c *EchoCustomCtx) GetPlatform() string {
	return strings.ToLower(c.Request().Header.Get(constants.HeaderPlatform))
}

// GetTimeFromAtQuery ...
func (c *EchoCustomCtx) GetTimeFromAtQuery() time.Time {
	return ptime.TimeParseISODate(c.QueryParam("fromAt"))
}

// GetTimeToAtQuery ...
func (c *EchoCustomCtx) GetTimeToAtQuery() time.Time {
	return ptime.TimeParseISODate(c.QueryParam("toAt"))
}
