package databasemongodb

import (
	"affiliate/internal/module/logger"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabase interface {
	Client() *mongo.Client
	Database() *mongo.Database
	Collection(model IModel) *mongo.Collection
	Find(ctx context.Context, model IModel, filter interface{}, opts ...*options.FindOptions) func(i interface{}) error
	CountByCondition(ctx context.Context, model IModel, filter interface{}) int64
	FindById(ctx context.Context, model IModel, id interface{}) error
	FindOne(ctx context.Context, model IModel, filter interface{}) error
	InsertOne(ctx context.Context, model IModel) error
	InsertMany(ctx context.Context, model IModel, payload []interface{}) error
	UpdateOne(ctx context.Context, model IModel, cond interface{}, update interface{}) error
	UpdateMany(ctx context.Context, model IModel, cond interface{}, update interface{}) error
	DeleteOne(ctx context.Context, model IModel, cond interface{}) error
	DeleteMany(ctx context.Context, model IModel, cond interface{}) error
}

type IModel interface {
	DbModelName() string
}

func NewDBShare(client *mongo.Client, nameDB string) IDatabase {
	return &dbShare{
		client: client,
		name:   nameDB,
	}
}

type dbShare struct {
	client *mongo.Client
	name   string
}

func (d dbShare) CountByCondition(ctx context.Context, model IModel, filter interface{}) int64 {
	total, _ := d.Collection(model).CountDocuments(ctx, filter)
	return total
}

func (d dbShare) Find(ctx context.Context, model IModel, filter interface{}, opts ...*options.FindOptions) func(i interface{}) error {
	cursor, err := d.Collection(model).Find(ctx, filter, opts...)
	return func(i interface{}) error {
		if err == nil {
			err = cursor.All(ctx, i)
		}
		if err != nil {
			return d.returnError(model, filter, "Find", err)
		}
		return d.returnError(model, filter, "Find", err)
	}
}

func (d dbShare) DeleteOne(ctx context.Context, model IModel, cond interface{}) error {
	_, err := d.Collection(model).DeleteOne(ctx, cond)
	return d.returnError(model, cond, "DeleteOne", err)
}

func (d dbShare) DeleteMany(ctx context.Context, model IModel, cond interface{}) error {
	_, err := d.Collection(model).DeleteMany(ctx, cond)
	return d.returnError(model, cond, "DeleteMany", err)
}

func (d dbShare) Collection(model IModel) *mongo.Collection {
	return d.Database().Collection(model.DbModelName())
}

func (d dbShare) Client() *mongo.Client {
	return d.client
}

func (d dbShare) Database() *mongo.Database {
	return d.client.Database(d.name)
}

func (d dbShare) InsertMany(ctx context.Context, model IModel, payload []interface{}) error {
	_, err := d.Collection(model).InsertMany(ctx, payload)
	return d.returnError(model, payload, "InsertMany", err)
}

func (d dbShare) FindById(ctx context.Context, model IModel, id interface{}) error {
	err := d.Collection(model).FindOne(ctx, bson.M{"_id": id}).Decode(model)
	if errors.Is(err, mongo.ErrNoDocuments) {
		err = nil
	}
	return d.returnError(model, id, "FindById", err)
}

func (d dbShare) FindOne(ctx context.Context, model IModel, filter interface{}) error {
	err := d.Collection(model).FindOne(ctx, filter).Decode(model)
	if errors.Is(err, mongo.ErrNoDocuments) {
		err = nil
	}
	return d.returnError(model, filter, "FindOne", err)
}

func (d dbShare) InsertOne(ctx context.Context, model IModel) error {
	_, err := d.Collection(model).InsertOne(ctx, model)
	return d.returnError(model, nil, "InsertOne", err)
}

func (d dbShare) UpdateMany(ctx context.Context, model IModel, cond interface{}, update interface{}) error {
	_, err := d.Collection(model).UpdateMany(ctx, cond, update)
	return d.returnError(model, cond, "UpdateMany", err)
}

func (d dbShare) UpdateOne(ctx context.Context, model IModel, cond interface{}, update interface{}) error {
	_, err := d.Collection(model).UpdateOne(ctx, cond, update)
	return d.returnError(model, cond, "UpdateOne", err)
}

func (d dbShare) returnError(model IModel, data interface{}, action string, err error) error {
	if err != nil {
		logger.Error("Error process with database", logger.LogData{
			Source:  "Database",
			Message: err.Error(),
		})
		err = fmt.Errorf("%w - %s %s.%s %v ", err, action, d.name, model.DbModelName(), data)
		return err
	}
	return nil
}
