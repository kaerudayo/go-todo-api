package domain

import "time"

type User struct {
	Id        int
	name      string
	email     int
	password  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
