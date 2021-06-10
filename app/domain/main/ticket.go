package domain

import "time"

type Ticket struct {
	Id          int       `db:"id, primarykey"`
	EpicId      int       `db:"epic_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
