package repository

import (
	domain "todo/app/domain/main"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Add(db *gorm.DB, ticket domain.Ticket) error
	Get(db *gorm.DB, id int) domain.Ticket
}
