package middleware

import (
	"affiliate/internal/constants"
	"net/http"
	"strings"

	"golang.org/x/text/language"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSConfig ...
func CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     constants.ListHeaderAllow,
		AllowCredentials: false,
		MaxAge:           600,
	})
}

// GetDeviceID ...
func GetDeviceID(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderDeviceID)
}

// GetAppVersion ...
func GetAppVersion(c echo.Context) string {
	return strings.ToLower(c.Request().Header.Get(constants.HeaderAppVersion))
}

// GetUserAgent ...
func GetUserAgent(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderUserAgent)
}

// GetFCMToken ...
func GetFCMToken(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderFCMToken)
}

// GetModel ...
func GetModel(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderModel)
}

// GetHeaderIP ...
func GetHeaderIP(c echo.Context) string {
	return strings.TrimSpace(c.RealIP())
}

// GetOsName ...
func GetOsName(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderOSName)
}

// GetOsVersion ...
func GetOsVersion(c echo.Context) string {
	return c.Request().Header.Get(constants.HeaderOSVersion)
}

// GetHeaderLanguage ...
func GetHeaderLanguage(c echo.Context) string {
	lang := c.Request().Header.Get(constants.HeaderAcceptLanguage)
	if lang != language.English.String() {
		lang = language.Vietnamese.String()
	}
	return lang
}

// GetAuthToken ...
func GetAuthToken(c echo.Context) string {
	return strings.Split(c.Request().Header.Get(constants.HeaderAuthorization), " ")[1]
}
