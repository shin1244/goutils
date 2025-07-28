package goutils

import "database/sql"

func OpenDB(dbType, host, port, user, pw, dbname string) (*sql.DB, error) {
	db, err := sql.Open(dbType, "host="+host+" port="+port+" user="+user+" password="+pw+" dbname="+dbname+" sslmode=disable")

	return db, err
}
