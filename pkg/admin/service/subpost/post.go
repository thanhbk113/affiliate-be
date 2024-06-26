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
func (p subPostImpl) Create(ctx context.Context, body request.SubPost) error {
	if err := checkParIdExist(ctx, body.ParID); err != nil {
		return err
	}
	subPostRaw := &modelmg.SubPostRaw{
		ID:        modelmg.NewAppID(),
		ParID:     util.ConvertStringToObjectID(body.ParID),
		Title:     body.Title,
		Image:     body.Image,
		Content:   body.Content,
		CreatedAt: util.TimeNow(),
		UpdatedAt: util.TimeNow(),
	}

	if err := daomongodb.SubPostDAO().GetShare().InsertOne(ctx, subPostRaw); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}

	return nil
}

func checkParIdExist(ctx context.Context, parID string) error {
	parPostRaw := new(modelmg.ParPostRaw)
	if err := daomongodb.ParPostDAO().GetShare().FindOne(ctx, parPostRaw, bson.M{"_id": util.ConvertStringToObjectID(parID)}); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}
	if parPostRaw.ID.IsZero() {
		return errors.New("par id not exist")
	}
	return nil
}
