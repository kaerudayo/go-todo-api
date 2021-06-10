package repository

import (
	"database/sql"
	"domain/udb"
)

type UserRepository interface {
	Get(DB *sql.DB, id int) error
	Insert(DB *sql.DB, user udb.Ticket) error
	Update(DB *sql.DB, user udb.Ticket) error
	Delete(DB *sql.DB, id int) error
}
