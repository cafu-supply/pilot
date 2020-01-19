package postgresql

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

func openConn() sqlbuilder.Database {
	connSettings := postgresql.ConnectionURL{
		User:     "postgres",
		Password: "",
		Host:     "",
		Socket:   "",
		Database: "cafu_pilot",
		Options:  nil,
	}

	conn, err := postgresql.Open(connSettings)

	if err != nil {
		panic("SHIT NO DB")
	}

	return conn
}

func getReadConn() sqlbuilder.Database {
	return openConn()
}

func getWriteConn() sqlbuilder.Database {
	return openConn()
}