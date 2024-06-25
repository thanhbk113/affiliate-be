package modelmg

import (
	databasemongodb "affiliate/internal/module/database/mongodb"
	"time"
)

type SessionDAO interface {
	GetShare() databasemongodb.IDatabase
}

// SessionUserRaw ...
type SessionRaw struct {
	ID        AppID     `bson:"_id"`
	Staff     AppID     `bson:"staff"`
	Token     string    `bson:"token"`
	ExpireAt  time.Time `bson:"expireAt"`
	CreatedAt time.Time `bson:"createdAt"`
}

// DbModelName ...
func (s *SessionRaw) DbModelName() string {
	return "sessions"
}
