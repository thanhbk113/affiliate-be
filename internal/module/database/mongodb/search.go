package databasemongodb

import (
	"affiliate/internal/format"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// GenerateQuerySearchString ...
func GenerateQuerySearchString(s string) bson.M {
	return bson.M{
		"$regex": bsonx.Regex(format.NonAccentVietnamese(s), "i"),
	}
}
