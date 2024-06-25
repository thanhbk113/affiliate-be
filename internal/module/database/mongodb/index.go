package databasemongodb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/x/bsonx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexDatabaseInterface ...
type IndexDatabaseInterface interface {
	// Indexes ...
	Indexes()
}

// indexDatabase ...
type indexDatabase struct{}

// IndexDatabase ...
func IndexDatabase() IndexDatabaseInterface {
	return indexDatabase{}
}

//
// METHOD PUBLIC
//

// Indexes ...
func (i indexDatabase) Indexes() {
}

//
// METHOD PRIVATE
//

// newIndex ...
func (i indexDatabase) newIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}

		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}

	return mongo.IndexModel{Keys: doc}
}

// newUniqIndex ...
func (i indexDatabase) newUniqIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}
		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}
	opt := options.Index().SetUnique(true)
	return mongo.IndexModel{Keys: doc, Options: opt}
}

// process ...
func (i indexDatabase) process(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute * 30)
	_, err := col.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}
