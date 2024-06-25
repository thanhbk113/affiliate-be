package serviceparpost

import (
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/request"
	"affiliate/pkg/admin/model/response"
	"context"
)

type ParPostInterface interface {
	Create(ctx context.Context, body request.ParPost) error
	GetList(ctx context.Context, query *mgquery.CommonQuery) (response.ParPostAllResponse, error)
}

type parPostImpl struct {
}

// Staff ...
func ParPost() ParPostInterface {
	return &parPostImpl{}
}
