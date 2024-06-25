package modelmg

import (
	databasemongodb "affiliate/internal/module/database/mongodb"
	"time"
)

type IntroduceRaw struct {
	ID           AppID     `bson:"_id,omitempty"`
	ParentNameId string    `bson:"parentNameId,omitempty"`
	Content      string    `bson:"content,omitempty"`
	CreatedAt    time.Time `bson:"createdAt,omitempty"`
	UpdatedAt    time.Time `bson:"updatedAt,omitempty"`
}

// DbModelName ...
func (r *IntroduceRaw) DbModelName() string {
	return "introduces"
}

type IntroduceDAO interface {
	GetShare() databasemongodb.IDatabase
}
