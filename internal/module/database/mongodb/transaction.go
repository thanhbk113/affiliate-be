package databasemongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// TransactionDatabaseInterface ...
type TransactionDatabaseInterface interface {
	// WithTransaction ...
	WithTransaction(ctx context.Context, fn func(sessionContext mongo.SessionContext) error) error
}

type transImpl struct{}

// TransactionDatabase ...
func TransactionDatabase() TransactionDatabaseInterface {
	return transImpl{}
}

// WithTransaction ...
func (t transImpl) WithTransaction(ctx context.Context, fn func(sessionContext mongo.SessionContext) error) error {
	client := db.Client()

	session, err := client.StartSession()
	if err != nil {
		return err
	}

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	rf := readpref.Primary()

	maxCommitTime := time.Minute
	txnOpts := options.Transaction().SetReadPreference(rf).
		SetWriteConcern(wc).SetReadConcern(rc).SetMaxCommitTime(&maxCommitTime)

	defer session.EndSession(ctx)
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		var errTransaction error
		if errTransaction = session.StartTransaction(txnOpts); errTransaction != nil {
			return errTransaction
		}

		// Handle func
		errTransaction = fn(sessionContext)
		if errTransaction != nil {
			return errTransaction
		}

		// Commit
		if errTransaction = session.CommitTransaction(sessionContext); errTransaction != nil {
			return errTransaction
		}
		return nil
	})

	if err != nil {
		// Rollback
		if abortErr := session.AbortTransaction(ctx); abortErr != nil {
			return abortErr
		}
		return err
	}
	return nil
}
