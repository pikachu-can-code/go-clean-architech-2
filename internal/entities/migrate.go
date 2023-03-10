package entities

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Permission{},
		&Role{},
		//&OtherSchemaNeedMigrate{},
		//&OtherTableNeedMigrate{},
		//&OtherEntityNeedMigrate{},
		// Add object to here, rebuild and run goose in folder database to migrate
	)
}
