package goutils

import "database/sql"

// OpenDB는 주어진 데이터베이스 유형, 호스트, 포트, 사용자 이름, 비밀번호 및 데이터베이스 이름을 사용하여 데이터베이스 연결을 엽니다.
func OpenDB(dbType, host, port, user, pw, dbname string) (*sql.DB, error) {
	db, err := sql.Open(dbType, "host="+host+" port="+port+" user="+user+" password="+pw+" dbname="+dbname+" sslmode=disable")

	return db, err
}
