package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(c *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB.Production.Username,
		c.DB.Production.Password,
		c.DB.Production.Host,
		c.DB.Production.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
