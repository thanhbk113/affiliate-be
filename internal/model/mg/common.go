package modelmg

import "go.mongodb.org/mongo-driver/bson/primitive"

// AppID ...
type AppID = primitive.ObjectID

// NewAppID ...
func NewAppID() AppID {
	return primitive.NewObjectID()
}
