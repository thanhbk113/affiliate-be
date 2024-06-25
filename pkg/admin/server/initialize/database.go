package initialize

import (
	"affiliate/internal/config"
	databasemongodb "affiliate/internal/module/database/mongodb"
)

func database() {
	cfg := config.GetENV().MongoDB
	_, err := databasemongodb.Connect(databasemongodb.Config{
		Host:   cfg.Host,
		DBName: cfg.DBName,
		Standalone: &databasemongodb.ConnectStandaloneOpts{
			AuthMechanism: cfg.Mechanism,
			AuthSource:    cfg.Source,
			Username:      cfg.Username,
			Password:      cfg.Password,
		},
	})
	if err != nil {
		panic(err)
	}
}
