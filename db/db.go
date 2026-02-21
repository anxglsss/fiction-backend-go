package db

import (
	"log"

	"fiction-turnament/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(databaseURL string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Tournament{},
		&models.TournamentMatch{},
	); err != nil {
		return err
	}

	log.Println("База данных подключена, таблицы созданы")
	return nil
}
