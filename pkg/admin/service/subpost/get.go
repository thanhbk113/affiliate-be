package servicesubpost

import (
	"affiliate/internal/locale"
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/response"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func (p subPostImpl) GetList(ctx context.Context, query *mgquery.CommonQuery, id modelmg.AppID) (response.SubPostAllResponse, error) {
	var (
		IModel   = new(modelmg.SubPostRaw)
		products = make([]modelmg.SubPostRaw, 0)
		res      = make([]response.SubPostRes, 0)
		data     = response.SubPostAllResponse{
			Total: 0,
			Limit: query.Limit,
		}
		cond = bson.M{}
	)

	if err := checkParID(ctx, id); err != nil {
		return data, err
	}

	cond["parId"] = id

	total := daomongodb.ParPostDAO().GetShare().CountByCondition(ctx, IModel, cond)
	data.Total = total
	data.Limit = query.Limit

	err := daomongodb.ParPostDAO().GetShare().Find(ctx, IModel, cond, query.GetFindOptsUsingPage())(&products)

	if err != nil {
		return data, err
	}

	for _, product := range products {
		productRes := p.getBriefWithoutContent(product)

		res = append(res, productRes)
	}

	data.List = res

	return data, nil
}

func (p subPostImpl) GetDetail(ctx context.Context, id modelmg.AppID) (response.SubPostRes, error) {
	var (
		product = new(modelmg.SubPostRaw)
	)
	err := daomongodb.SubPostDAO().GetShare().FindOne(ctx, product, bson.M{"_id": id})
	if err != nil {
		return response.SubPostRes{}, err
	}
	return p.getBrief(*product), nil
}

func (s subPostImpl) getBrief(product modelmg.SubPostRaw) response.SubPostRes {
	return response.SubPostRes{
		ID:        product.ID.Hex(),
		Title:     product.Title,
		Content:   product.Content,
		ParID:     product.ParID.Hex(),
		Image:     product.Image,
		CreatedAt: util.TimeISO(product.CreatedAt),
		UpdatedAt: util.TimeISO(product.UpdatedAt),
	}
}

func (s subPostImpl) getBriefWithoutContent(product modelmg.SubPostRaw) response.SubPostRes {
	return response.SubPostRes{
		ID:        product.ID.Hex(),
		Title:     product.Title,
		ParID:     product.ParID.Hex(),
		Image:     product.Image,
		CreatedAt: util.TimeISO(product.CreatedAt),
		UpdatedAt: util.TimeISO(product.UpdatedAt),
	}
}

func checkParID(ctx context.Context, id modelmg.AppID) error {
	parPostRaw := new(modelmg.ParPostRaw)

	if err := daomongodb.ParPostDAO().GetShare().FindOne(ctx, parPostRaw, bson.M{"_id": id}); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}
	if parPostRaw.ID.IsZero() {
		return errors.New("par id not exist")
	}
	return nil
}

func checkSubID(ctx context.Context, id modelmg.AppID) error {
	subPostRaw := new(modelmg.SubPostRaw)

	if err := daomongodb.SubPostDAO().GetShare().FindOne(ctx, subPostRaw, bson.M{"_id": id}); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}
	if subPostRaw.ID.IsZero() {
		return errors.New("sub id not exist")
	}
	return nil
}
