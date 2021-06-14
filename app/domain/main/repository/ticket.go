package repository

import (
	domain "todo/app/domain/main"

	"github.com/go-gorp/gorp"
)

type TicketRepository interface {
	Add(db *gorp.DbMap, ticket domain.Ticket) error
	Get(db *gorp.DbMap, id int) (*domain.Ticket, error)
}
