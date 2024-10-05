package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/venedial/taskops/models"
	"github.com/venedial/taskops/storage"
)

type TaskStorer interface {
	Create(ctx context.Context, dto *models.TaskDTO) (*models.Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.Task, error)
	FindAll(ctx context.Context) ([]models.Task, error)
	UpdateState(ctx context.Context, id uuid.UUID, state models.TaskState) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskStore struct {
	storage *storage.TaskOpsStorage
}

func NewTaskStore(s *storage.TaskOpsStorage) *TaskStore {
	return &TaskStore{
		storage: s,
	}
}

// Create inserts a new task into the database.
func (ts *TaskStore) Create(ctx context.Context, dto *models.TaskDTO) (*models.Task, error) {
	task := &models.Task{
		Name:  dto.Name,
		State: models.TaskState("PENDING"),
	}

	if err := ts.storage.DB.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// FindByID retrieves a task by its ID.
func (ts *TaskStore) FindByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var task models.Task
	if err := ts.storage.DB.WithContext(ctx).First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

// FindAll retrieves all tasks from the database.
func (ts *TaskStore) FindAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	if err := ts.storage.DB.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// UpdateState updates the state of a task by its ID.
func (ts *TaskStore) UpdateState(ctx context.Context, id uuid.UUID, state models.TaskState) error {
	result := ts.storage.DB.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Update("state", state)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Delete removes a task from the database by its ID.
func (ts *TaskStore) Delete(ctx context.Context, id uuid.UUID) error {
	result := ts.storage.DB.WithContext(ctx).Delete(&models.Task{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
