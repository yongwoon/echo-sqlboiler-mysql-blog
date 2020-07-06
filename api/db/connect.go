package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// InitMysql initialize database connection
func Connect() (*sql.DB, error) {
	DBMS := "mysql"
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("DATABASE_SCHEMA")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"

	db, err := sql.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}

	return db, err
}
