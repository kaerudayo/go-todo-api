package repository

import (
	"database/sql"
	"domain/main"
)

type TicketRepository interface {
	Get(DB *sql.DB, id int) error
	Insert(DB *sql.DB, ticket main.Ticket) error
}
