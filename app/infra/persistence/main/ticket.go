package ticket

import (
	"fmt"
	domain "todo/app/domain/main"
	"todo/app/domain/main/repository"

	"gorm.io/gorm"
)

type tikcetPersistence struct{}

func NewTicketPersistence() repository.TicketRepository {
	return &tikcetPersistence{}
}

/**
 * method
 */
func (tp tikcetPersistence) Add(db *gorm.DB, ticket domain.Ticket) error {
	err := db.Create(&ticket)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (tp tikcetPersistence) Get(db *gorm.DB, id int) domain.Ticket {
	var data domain.Ticket
	db.First(&data, 1)
	return data
}
