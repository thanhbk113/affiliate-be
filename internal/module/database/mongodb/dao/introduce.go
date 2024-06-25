package daomongodb

import (
	modelmg "affiliate/internal/model/mg"
	databasemongodb "affiliate/internal/module/database/mongodb"
)

type introduceDAO struct {
	DbShare databasemongodb.IDatabase
}

// GetShare ...
func (r *introduceDAO) GetShare() databasemongodb.IDatabase {
	return r.DbShare
}

// IntroduceDAO ...
func IntroduceDAO() modelmg.IntroduceDAO {
	return &introduceDAO{
		DbShare: databasemongodb.GetDBShare(),
	}
}
