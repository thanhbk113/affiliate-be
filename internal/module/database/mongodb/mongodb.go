package databasemongodb

import (
	"context"
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config ...
type Config struct {
	Host       string
	DBName     string
	Monitor    *event.CommandMonitor
	Standalone *ConnectStandaloneOpts
}

// ConnectStandaloneOpts ...
type ConnectStandaloneOpts struct {
	AuthMechanism string
	AuthSource    string
	Username      string
	Password      string
}

var (
	db  *mongo.Database
	iDB IDatabase
)

// Connect to mongo server
func Connect(cfg Config) (*mongo.Database, error) {
	connectOptions := options.ClientOptions{}
	opts := cfg.Standalone
	// Set auth if existed
	if opts != nil && opts.Username != "" && opts.Password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: opts.AuthMechanism,
			AuthSource:    opts.AuthSource,
			Username:      opts.Username,
			Password:      opts.Password,
		}
	}
	if cfg.Monitor != nil {
		connectOptions.SetMonitor(cfg.Monitor)
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(cfg.Host))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", cfg.Host, err)
		return nil, err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + cfg.Host + " --- DB: " + cfg.DBName))

	// Set data
	db = client.Database(cfg.DBName)
	iDB = NewDBShare(db.Client(), cfg.DBName)
	return db, nil
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}

// GetDBShare ...
func GetDBShare() IDatabase {
	return iDB
}

func getReadPref(mode string) *readpref.ReadPref {
	m, err := readpref.ModeFromString(mode)
	if err != nil {
		m = readpref.SecondaryPreferredMode
	}
	readPref, err := readpref.New(m)
	if err != nil {
		fmt.Println("mongodb.getReadPref err: ", err, m)
	}
	return readPref
}
