package daomongodb

import (
	modelmg "affiliate/internal/model/mg"
	databasemongodb "affiliate/internal/module/database/mongodb"
)

type sessionDAO struct {
	DbShare databasemongodb.IDatabase
}

// GetShare ...
func (r *sessionDAO) GetShare() databasemongodb.IDatabase {
	return r.DbShare
}

// SessionDAO ...
func SessionDAO() modelmg.SessionDAO {
	return &sessionDAO{
		DbShare: databasemongodb.GetDBShare(),
	}
}
