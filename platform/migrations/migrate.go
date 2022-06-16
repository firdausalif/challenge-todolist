package migrations

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	log.Println("Initiating migration...")
	err := db.Migrator().AutoMigrate()
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
