package daomongodb

import (
	modelmg "affiliate/internal/model/mg"
	databasemongodb "affiliate/internal/module/database/mongodb"
)

type parPostDAO struct {
	DbShare databasemongodb.IDatabase
}

// GetShare ...
func (r *parPostDAO) GetShare() databasemongodb.IDatabase {
	return r.DbShare
}

// SubPostDAO ...
func ParPostDAO() modelmg.ParPostDAO {
	return &parPostDAO{
		DbShare: databasemongodb.GetDBShare(),
	}
}
