package migrations

import (
	"github.com/firdausalif/challenge-todolist/app/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	log.Println("Initiating migration...")
	err := db.Migrator().AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
