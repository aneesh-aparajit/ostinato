package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskDTO struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	ExecutionJs string     `json:"executionJs"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (t *TaskDTO) GetPGTypeID() pgtype.UUID {
	if t.ID == nil {
		id := uuid.New()
		t.ID = &id
	}
	return pgtype.UUID{
		Bytes: *t.ID,
		Valid: true,
	}
}
