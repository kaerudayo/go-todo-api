package repository

import (
	"database/sql"
	"domain/main"
)

type EpicRepository interface {
	Get(DB *sql.DB, id int) error
	Insert(DB *sql.DB, epic main.Epic) error
}
