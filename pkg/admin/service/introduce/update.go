package servicesintroduce

import (
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/pkg/admin/model/request"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *introduceImpl) Update(ctx context.Context, body request.IntroduceBodyUpdate) error {

	update := &modelmg.IntroduceRaw{
		ParentNameId: s.User.NameID,
		Content:      body.Content,
		UpdatedAt:    util.TimeNow(),
	}

	find := bson.M{"parentNameId": s.User.NameID}

	if err := daomongodb.IntroduceDAO().GetShare().UpdateOne(ctx, update, find, bson.M{"$set": update}); err != nil {
		return err
	}
	return nil
}
