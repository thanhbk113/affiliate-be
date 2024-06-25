package serviceparpost

import (
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/response"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (p parPostImpl) GetList(ctx context.Context, query *mgquery.CommonQuery) (response.ParPostAllResponse, error) {
	var (
		IModel   = new(modelmg.ParPostRaw)
		products = make([]modelmg.ParPostRaw, 0)
		res      = make([]response.ParPostRes, 0)
		data     = response.ParPostAllResponse{
			Total: 0,
			Limit: query.Limit,
		}
		cond = bson.M{}
	)

	total := daomongodb.ParPostDAO().GetShare().CountByCondition(ctx, IModel, cond)
	data.Total = total
	data.Limit = query.Limit

	err := daomongodb.ParPostDAO().GetShare().Find(ctx, IModel, cond, query.GetFindOptsUsingPage())(&products)

	if err != nil {
		return data, err
	}

	for _, product := range products {
		productRes := p.getBrief(product)

		res = append(res, productRes)
	}

	data.List = res

	return data, nil
}

func (s parPostImpl) getBrief(product modelmg.ParPostRaw) response.ParPostRes {
	return response.ParPostRes{
		ID:        product.ID.Hex(),
		Name:      product.Name,
		CreatedAt: util.TimeISO(product.CreatedAt),
		UpdatedAt: util.TimeISO(product.UpdatedAt),
	}
}
