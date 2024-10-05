package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/venedial/taskops/config"
	"github.com/venedial/taskops/logger"
	"github.com/venedial/taskops/models"
	"github.com/venedial/taskops/storage"
	"github.com/venedial/taskops/store"
)

// USED ONLY FOR DB INITIALIZATION REMOVE AFTER ISSUE-4 IS DONE
func main() {
	ctx := context.Background()

	// Load config
	conf := config.Conf()

	// Setup default logger
	logger.Setup(conf)
	defer logger.Cleanup()

	// Setup main storage or panic
	toStorage, err := storage.NewTaskOpsStorage(conf)
	if err != nil {
		slog.Error("Unable to initialize store!", "error", err)
		panic(fmt.Errorf("Unable to initialize store. Error: %v", err))
	}
	toStorage.Migrate()

	taskStore := store.NewTaskStore(toStorage)

	result, err := taskStore.Create(
		ctx,
		&models.TaskDTO{Name: "test", State: models.TaskState("FAILED")},
	)
	if err != nil {
		slog.Error("Something went wrong!", err)
	}

	id, err := uuid.Parse("1e53649f-89fd-401a-b5ce-ff7d194633a6")
	if err != nil {
		slog.Error("Something went wrong", err)
	}

	result, err = taskStore.FindByID(ctx, id)
	if err != nil {
		fmt.Println("Something went wrong!", err)
	}

	slog.Info("task", result)
	slog.Info("App running!")
}
