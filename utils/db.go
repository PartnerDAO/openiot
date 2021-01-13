package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	DB, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/openiot20210109")
	if err != nil {
		panic(err.Error())
	}
}
