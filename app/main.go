package main

import (
	"fmt"
	domain "todo/app/domain/main"
	infrastructure "todo/app/infra"
	"todo/app/infra/persistence/ticket"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := infrastructure.InitDb()
	defer db.Db.Close()

	u := ticket.NewTicketPersistence()
	data := domain.Ticket{
		EpicId:      1,
		Title:       "aaa",
		Description: "bbbbb",
	}
	v := u.Add(db, data)
	res, err := u.Get(db, 1)
	if err != nil {
		fmt.Println(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data1": v,
			"data2": res,
		})
	})
	r.Run(":3000")
}
