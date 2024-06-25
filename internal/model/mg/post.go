package modelmg

import (
	databasemongodb "affiliate/internal/module/database/mongodb"
	"time"
)

type ParPostRaw struct {
	ID        AppID     `bson:"_id"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type SubPostRaw struct {
	ID        AppID     `bson:"_id"`
	ParID     AppID     `bson:"parId"`
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (r *ParPostRaw) DbModelName() string {
	return "parPosts"
}

type ParPostDAO interface {
	GetShare() databasemongodb.IDatabase
}

func (r *SubPostRaw) DbModelName() string {
	return "subPosts"
}

type SubPostDAO interface {
	GetShare() databasemongodb.IDatabase
}
