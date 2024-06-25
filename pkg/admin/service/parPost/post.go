package serviceparpost

import (
	"affiliate/internal/locale"
	modelmg "affiliate/internal/model/mg"
	daomongodb "affiliate/internal/module/database/mongodb/dao"
	"affiliate/internal/util"
	"affiliate/pkg/admin/model/request"
	"context"
	"errors"
)

// Create ...
func (p parPostImpl) Create(ctx context.Context, body request.ParPost) error {
	parPostRaw := &modelmg.ParPostRaw{
		ID:        modelmg.NewAppID(),
		Name:      body.Name,
		CreatedAt: util.TimeNow(),
		UpdatedAt: util.TimeNow(),
	}

	if err := daomongodb.ParPostDAO().GetShare().InsertOne(ctx, parPostRaw); err != nil {
		return errors.New(locale.CommonKeyBadRequest)
	}

	return nil
}
