package servicesintroduce

import (
	"affiliate/internal/middleware"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/request"
	"affiliate/pkg/admin/model/response"
	"context"
)

type IntroduceInterface interface {
	Update(ctx context.Context, body request.IntroduceBodyUpdate) error
	GetDetail(ctx context.Context, q *mgquery.CommonQuery) (response.IntroduceResponse, error)
}

type introduceImpl struct {
	User *middleware.User
}

// Staff ...
func Introduce(user *middleware.User) IntroduceInterface {
	return &introduceImpl{
		User: user,
	}
}
