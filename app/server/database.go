package server

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB() (err error) {
	var dataSourceName = "root:1234@/users"
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
	}

	err = Db.Ping()
	return
}
