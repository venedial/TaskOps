package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Name  string    `gorm:"size:255"`
	State TaskState `gorm:"type:taskState;default:'PENDING'"`
}

func (t *Task) ToDto() *TaskDTO {
	return &TaskDTO{
		ID:    t.ID,
		Name:  t.Name,
		State: t.State,
	}
}

type TaskDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	State TaskState `json:"state"`
}

type TaskState string

const (
	PENDING    TaskState = "PENDING"
	ENQUEUED   TaskState = "ENQUEUED"
	PROCESSING TaskState = "PROCESSING"
	SUCCEEDED  TaskState = "SUCCEEDED"
	DELETED    TaskState = "DELETED"
	FAILED     TaskState = "FAILED"
	TIMEOUT    TaskState = "TIMEOUT"
)

func (p *TaskState) Scan(value interface{}) error {
	*p = TaskState(value.(string))
	return nil
}

func (p TaskState) Value() (driver.Value, error) {
	return string(p), nil
}
