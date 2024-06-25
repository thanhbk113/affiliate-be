package servicesubpost

import (
	"affiliate/internal/locale"
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/pkg/admin/model/request"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// Create ...
func (p subPostImpl) Update(ctx context.Context, body request.SubPostUpdate, id modelmg.AppID) error {
	if err := checkParIdExist(ctx, id.Hex()); err != nil {
		return err
	}
	IModel := new(modelmg.SubPostRaw)
	update := bson.M{"$set": bson.M{"content": body.Content, "updatedAt": util.TimeNow()}}
	if err := daomongodb.SubPostDAO().GetShare().UpdateOne(ctx, IModel, bson.M{"_id": id}, update); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}

	return nil
}
