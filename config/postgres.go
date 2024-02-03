package config

import (
	"github.com/fymorGod/go-alura-apirest/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresDb() (*gorm.DB, error) {
	dsn := "host=localhost user=fymorpgadmin password=fymorpgpassword# dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}
	db.AutoMigrate(&schemas.Aluno{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
