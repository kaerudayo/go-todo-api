package main

import (
	domain "app/app/domain/udb"
	"domain/main/repository"

	"github.com/go-gorp/gorp"
)

type tikcetPersistence struct{}

func NewTicketPersistence() repository.TicketRepository {
	return &ticketPersistence{}
}

/**
 * method
 */
func (tp tikcetPersistence) Insert(db gorp.SqlExecutor, ticket domain.Ticket) error {
	stmt, err := db.Inset(ticket)
	if err != nil {
		return err
	}
	return err
}

func (tp tikectPersistence) Get(db gorp.SqlExecutor, id int) (domain.Ticket, error) {
	t, err := dbmap.Get(domain.Ticket{}, 1)
	if err != nil {
		return err
	}
	return t
}
