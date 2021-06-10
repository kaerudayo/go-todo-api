package domain

import "time"

type Epic struct {
	Id        int       `db:"id, primarykey"`
	Title     string    `db:"title"`
	Color     int       `db:"color"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
