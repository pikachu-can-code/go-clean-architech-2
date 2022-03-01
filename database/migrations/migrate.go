package migrations

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(connection string) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db)
}
