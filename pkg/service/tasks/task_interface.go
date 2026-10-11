package tasks

import (
	"context"

	"github.com/aneesh-aparajit/ostinato/pkg/dto"
)

type ITaskService interface {
	CreateTask(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error)
	UpdateTask(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error)
	DeleteTask(ctx context.Context, dto *dto.TaskDTO) error
	GetTaskById(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error)
}
