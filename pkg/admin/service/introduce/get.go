package servicesintroduce

import (
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/internal/util/mgquery"
	"affiliate/pkg/admin/model/response"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s introduceImpl) GetDetail(ctx context.Context, q *mgquery.CommonQuery) (response.IntroduceResponse, error) {
	var (
		introduce = new(modelmg.IntroduceRaw)
		cond      = bson.M{}
	)

	q.AssignParentNameId(&cond)

	if err := daomongodb.IntroduceDAO().GetShare().FindOne(ctx, introduce, cond); err != nil {
		return response.IntroduceResponse{}, err
	}

	if introduce.ID.IsZero() {
		return response.IntroduceResponse{}, nil
	}

	data := response.IntroduceResponse{
		ID:        introduce.ID.Hex(),
		Content:   introduce.Content,
		CreatedAt: util.TimeISO(introduce.CreatedAt),
		UpdatedAt: util.TimeISO(introduce.UpdatedAt),
	}

	return data, nil
}
