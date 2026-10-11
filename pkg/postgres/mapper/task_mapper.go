package mapper

import (
	"github.com/aneesh-aparajit/ostinato/pkg/dto"
	"github.com/aneesh-aparajit/ostinato/pkg/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapTask(dto *dto.TaskDTO) *postgres.Task {
	if dto == nil {
		return nil
	}

	var id pgtype.UUID
	if dto.ID != nil {
		id = pgtype.UUID{Bytes: *dto.ID, Valid: true}
	}

	return &postgres.Task{
		ID:          id,
		Name:        dto.Name,
		ExecutionJs: dto.ExecutionJs,
		CreatedAt:   pgtype.Timestamptz{Time: dto.CreatedAt, Valid: !dto.CreatedAt.IsZero()},
		UpdatedAt:   pgtype.Timestamptz{Time: dto.UpdatedAt, Valid: !dto.UpdatedAt.IsZero()},
	}
}

func MapTaskDTO(task *postgres.Task) *dto.TaskDTO {
	if task == nil {
		return nil
	}

	var id *uuid.UUID
	if task.ID.Valid {
		u := uuid.UUID(task.ID.Bytes)
		id = &u
	}

	return &dto.TaskDTO{
		ID:          id,
		Name:        task.Name,
		ExecutionJs: task.ExecutionJs,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}
}
