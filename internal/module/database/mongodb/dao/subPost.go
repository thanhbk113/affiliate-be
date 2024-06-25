package daomongodb

import (
	modelmg "affiliate/internal/model/mg"
	databasemongodb "affiliate/internal/module/database/mongodb"
)

type subPostDAO struct {
	DbShare databasemongodb.IDatabase
}

// GetShare ...
func (r *subPostDAO) GetShare() databasemongodb.IDatabase {
	return r.DbShare
}

// SubPostDAO ...
func SubPostDAO() modelmg.IntroduceDAO {
	return &subPostDAO{
		DbShare: databasemongodb.GetDBShare(),
	}
}
