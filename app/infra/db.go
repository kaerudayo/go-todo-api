package infrastructure

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "tcp:mysql-container:3307*todo/root/todo")
	if err != nil {
		log.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	return dbmap
}
