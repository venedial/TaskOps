package db

import (
	"golang.org/x/exp/slog"

	"taskops/internal/task"
)

var EnableUuidOsspExtensionQuery string = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

func RunMigration() {
	if DB == nil {
		slog.Error("Can't run migration on non initialized db!")
		return
	}

	// Enable only if Postgres version is < 14
	// DB.Exec(EnableUuidOsspExtensionQuery)
	DB.Exec(task.CreateTypeQuery)
	DB.AutoMigrate(&task.Task{})
}
