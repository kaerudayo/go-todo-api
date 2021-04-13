package reposiroy

import "app/domain/model"
import (
  "fmt"
  "log"
	_ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
)

func GetAll() {
  db, err := sqlx.Open("mysql","todo:todo@(mysql-container:3306)/todo")
	if err != nil {
		log.Fatalln(err)
	}
  if err := db.Ping();err != nil {
    log.Fatal(err)
  }

  tickets := []model.Ticket{}
  ticket  := model.Ticket{}
  r,e := db.Queryx("select * from ticket")
  for r.Next() {
    err := r.StructScan(&ticket)
    if err != nil {
        log.Fatal(err)
    }
    tickets = append(tickets, ticket)
  }
  fmt.Println(r)
  fmt.Println(e)
  fmt.Println(tickets)

  db.Close()
}

