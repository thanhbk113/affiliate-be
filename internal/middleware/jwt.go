package middleware

import (
	"affiliate/internal/constants"
	"affiliate/internal/errorresponse"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWT ...
func JWT(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(secret),
		Skipper: func(c echo.Context) bool {
			authHeader := c.Request().Header.Get(constants.HeaderAuthorization)
			// return true to skip middleware
			return authHeader == "" || authHeader == "Bearer" || authHeader == "Bearer "
		},
	})
}

type User struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	NameID string `json:"nameID"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
	Type   string `json:"type"`
}

// GetCurrentUserByToken ...
func GetCurrentUserByToken(token *jwt.MapClaims) *User {
	if token == nil {
		return nil
	}

	var (
		user = new(User)
	)

	if (*token)["_id"] != "" {
		user.ID = (*token)["_id"].(string)
	}

	if (*token)["name"] != "" {
		name, ok := (*token)["name"]
		if ok {
			user.Name = name.(string)
		}
	}

	if (*token)["email"] != "" {
		email, ok := (*token)["email"]
		if ok {
			user.Email = email.(string)
		}
	}

	if (*token)["exp"] != "" {
		exp, ok := (*token)["exp"]
		if ok {
			user.Exp = int64(exp.(float64))
		}
	}

	if (*token)["type"] != "" {
		exp, ok := (*token)["type"]
		if ok {
			user.Type = exp.(string)
		}
	}

	if (*token)["nameID"] != "" {
		nameID, ok := (*token)["nameID"]
		if ok {
			user.NameID = nameID.(string)
		}
	}

	return user
}

// DecodeToken ...
func DecodeToken(tokenString string, secretKey string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		if time.Now().Unix() > int64(exp) {
			return nil, errors.New(errorresponse.CommonKeyTokenAuthenticationExpired)
		}
		return &claims, nil
	}

	return nil, errors.New(errorresponse.CommonKeyTokenAuthenticationExpired)
}
