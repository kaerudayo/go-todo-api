package main

import (
	domain "todo/app/domain/main"
	infrastructure "todo/app/infra"
	persistence "todo/app/infra/persistence/main"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	conf := infrastructure.NewConfig()
	db := infrastructure.Init(conf)

	u := persistence.NewTicketPersistence()
	data := domain.Ticket{
		EpicId:      1,
		Title:       "aaa",
		Description: "bbbbb",
	}
	v := u.Add(db, data)
	res := u.Get(db, 1)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data1": v,
			"data2": res,
		})
	})
	r.Run(conf.Routing.Port)
}
