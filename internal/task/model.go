package task

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Name  string    `gorm:"size:255"`
	State State     `gorm:"type:taskState;default:'PENDING'"`
}

func (t *Task) ToDto() *DTO {
	return &DTO{
		ID:    t.ID,
		Name:  t.Name,
		State: t.State,
	}
}

type DTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	State State     `json:"state"`
}

type State string

var (
	PENDING    State = "PENDING"
	ENQUEUED   State = "ENQUEUED"
	PROCESSING State = "PROCESSING"
	SUCCEEDED  State = "SUCCEEDED"
	DELETED    State = "DELETED"
	FAILED     State = "FAILED"
	TIMEOUT    State = "TIMEOUT"
)

// Make sure that all enum variants are added here
var CreateTypeQuery string = `
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

func (p *State) Scan(value interface{}) error {
	*p = State(value.(string))
	return nil
}

func (p State) Value() (driver.Value, error) {
	return string(p), nil
}
