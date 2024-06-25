package servicesubpost

import (
	modelmg "affiliate/internal/model/mg"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/request"
	"affiliate/pkg/admin/model/response"
	"context"
)

type SubPostInterface interface {
	Create(ctx context.Context, body request.SubPost) error
	GetList(ctx context.Context, query *mgquery.CommonQuery, id modelmg.AppID) (response.SubPostAllResponse, error)
}

type subPostImpl struct {
}

// SubPost ...
func SubPost() SubPostInterface {
	return &subPostImpl{}
}
