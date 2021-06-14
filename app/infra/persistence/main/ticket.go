package ticket

import (
	domain "todo/app/domain/main"
	"todo/app/domain/main/repository"

	"github.com/go-gorp/gorp"
)

type tikcetPersistence struct{}

func NewTicketPersistence() repository.TicketRepository {
	return &tikcetPersistence{}
}

/**
 * method
 */
func (tp tikcetPersistence) Add(db *gorp.DbMap, ticket domain.Ticket) error {
	err := db.Insert(ticket)
	if err != nil {
		return err
	}
	return err
}

func (tp tikcetPersistence) Get(db *gorp.DbMap, id int) (*domain.Ticket, error) {
	var data domain.Ticket
	_, err := db.Get(data, 1)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
