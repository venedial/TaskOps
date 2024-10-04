package db

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"taskops/config"
)

var DB *gorm.DB

func Connect(conf *config.Configuration) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		conf.Pgdb.Host,
		conf.Pgdb.Port,
		conf.Pgdb.Database,
		conf.Pgdb.Username,
		conf.Pgdb.Password,
		conf.Pgdb.Sslmode,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	slog.Info("Database sucessfully initialized!", "database", conf.Pgdb.Database, "host", conf.Pgdb.Host)
}
