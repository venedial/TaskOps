package storage

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/venedial/taskops/config"
	"github.com/venedial/taskops/models"
)

type TaskOpsStorage struct {
	DB *gorm.DB
}

func NewTaskOpsStorage(conf *config.Configuration) (*TaskOpsStorage, error) {
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	slog.Info("Database sucessfully initialized!", "database", conf.Pgdb.Database, "host", conf.Pgdb.Host)

	return &TaskOpsStorage{DB: db}, nil
}

// Run only if Postgres version is < 14
const enableUuidOsspExtensionQuery string = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

// Make sure that all enum variants are added here
const createTaskTypeQuery string = `
  DO $$
  BEGIN
      IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'taskstate') THEN
          CREATE TYPE taskState AS ENUM (
              'PENDING', 
              'ENQUEUED', 
              'PROCESSING', 
              'SUCCEEDED', 
              'DELETED', 
              'FAILED', 
              'TIMEOUT'
          );
      END IF;
  END $$;
`

func (s *TaskOpsStorage) Migrate() {
	s.DB.Exec(createTaskTypeQuery)
	s.DB.AutoMigrate(&models.Task{})
}
