package db

import (
	"log"
	"strings"

	"fiction-turnament/models"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(databaseURL string) error {
	var err error
	if strings.HasPrefix(databaseURL, "postgres") {
		DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		log.Println("Подключено к PostgreSQL")
	} else {
		DB, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
		log.Println("Подключено к SQLite")
	}
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
