package task

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbRepository struct {
	db *gorm.DB
}

func NewDbRepository(db *gorm.DB) TaskRepository {
	return &DbRepository{
		db: db,
	}
}

func (r *DbRepository) Create(ctx context.Context, dto *DTO) (*Task, error) {
	task := &Task{
		Name:  dto.Name,
		State: PENDING,
	}

	if err := r.db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *DbRepository) FindByID(ctx context.Context, id uuid.UUID) (*Task, error) {
	var task Task
	if err := r.db.WithContext(ctx).First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *DbRepository) FindAll(ctx context.Context) ([]Task, error) {
	var tasks []Task
	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *DbRepository) UpdateState(ctx context.Context, id uuid.UUID, state State) error {
	result := r.db.WithContext(ctx).Model(&Task{}).Where("id = ?", id).Update("state", state)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *DbRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
