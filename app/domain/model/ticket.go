package model

import (
  "time"
)

// declare table
type Ticket struct {
  Id        int64     `db:"id"`
  Title     string    `db:"title"`
  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
