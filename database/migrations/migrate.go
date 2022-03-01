package migrations

import (
	"log"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(connection string) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	entities.Migrate(db)
}
