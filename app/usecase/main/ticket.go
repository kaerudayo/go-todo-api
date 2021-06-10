package main

import (
	domain "app/app/domain/main"
	"app/app/domain/main/repository"

	"github.com/go-gorp/gorp"
)

type TicketUseCase interface {
	Get(db gorp.SqlExecutor, id string) (domain.Ticket, error)
	Insert(db gorp.SqlExecutor, ticket domain.Ticket) error
}

type ticketUseCase struct {
	ticketRepository repository.TicketRepository
}

func NewTicketUseCase(tp repository.TicketRepository) TicketUseCase {
	return &ticketUseCase{
		ticketRepository: tp,
	}
}

func (tu TicketUsecase) Get(db gorp.SqlExecutor, id string) (domain.Ticket, error) {
	tikcet, err := tu.ticketRepository.Get(db, id)
	if err != nil {
		return nil, err
	}
	return tikcet, nil
}

func (tu TicketUsecase) INsert(db gorp.SqlExecutor, ticket domain.Ticket) error {
	tikcet, err := tu.ticketRepository.Insert(db, ticket)
	if err != nil {
		return nil, err
	}
	return user, nil
}
