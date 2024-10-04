package task

import (
	"context"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(ctx context.Context, dto *DTO) (*Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Task, error)
	FindAll(ctx context.Context) ([]Task, error)
	UpdateState(ctx context.Context, id uuid.UUID, state State) error
	Delete(ctx context.Context, id uuid.UUID) error
}
