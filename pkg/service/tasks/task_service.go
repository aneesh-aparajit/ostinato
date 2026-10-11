package tasks

import (
	"context"
	"time"

	"github.com/aneesh-aparajit/ostinato/pkg/dto"
	"github.com/aneesh-aparajit/ostinato/pkg/logger"
	"github.com/aneesh-aparajit/ostinato/pkg/postgres"
	"github.com/aneesh-aparajit/ostinato/pkg/postgres/mapper"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type TaskService struct {
	db         *pgxpool.Pool
	repository *postgres.Queries
}

func NewTaskService(db *pgxpool.Pool) *TaskService {
	return &TaskService{
		db:         db,
		repository: postgres.New(db),
	}
}

func (t *TaskService) GetTaskById(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error) {
	task, err := t.repository.GetTaskById(ctx, dto.GetPGTypeID())
	if err != nil {
		logger.FromContext(ctx).Error("unable to get task by id", zap.Error(err))
		return nil, err
	}
	dto = mapper.MapTaskDTO(&task)
	return dto, nil
}

func (t *TaskService) CreateTask(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		logger.FromContext(ctx).Error("unable to start transaction", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback(ctx)
	rtx := t.repository.WithTx(tx)
	task, err := rtx.CreateTask(ctx, postgres.CreateTaskParams{
		ID:          dto.GetPGTypeID(),
		Name:        dto.Name,
		ExecutionJs: dto.ExecutionJs,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now().UTC()},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now().UTC()},
	})
	if err != nil {
		logger.FromContext(ctx).Error("unable to create task", zap.Error(err))
		return nil, err
	}
	dto = mapper.MapTaskDTO(&task)
	if err := tx.Commit(ctx); err != nil {
		logger.FromContext(ctx).Error("unable to commit transaction", zap.Error(err))
		return nil, err
	}
	return dto, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, dto *dto.TaskDTO) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		logger.FromContext(ctx).Error("unable to start transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback(ctx)
	rtx := t.repository.WithTx(tx)
	if err := rtx.DeleteById(ctx, dto.GetPGTypeID()); err != nil {
		logger.FromContext(ctx).Error("unable to delete task", zap.Error(err))
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		logger.FromContext(ctx).Error("unable to commit transaction", zap.Error(err))
		return err
	}
	return nil
}

func (t *TaskService) UpdateTask(ctx context.Context, dto *dto.TaskDTO) (*dto.TaskDTO, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		logger.FromContext(ctx).Error("unable to start transaction", zap.Error(err))
		return nil, err
	}
	defer tx.Rollback(ctx)
	rtx := t.repository.WithTx(tx)
	task, err := rtx.UpdateTask(ctx, postgres.UpdateTaskParams{
		ID:          dto.GetPGTypeID(),
		Name:        dto.Name,
		ExecutionJs: dto.ExecutionJs,
	})
	if err != nil {
		logger.FromContext(ctx).Error("unable to update task", zap.Error(err))
		return nil, err
	}
	dto = mapper.MapTaskDTO(&task)
	if err := tx.Commit(ctx); err != nil {
		logger.FromContext(ctx).Error("unable to commit transaction", zap.Error(err))
		return nil, err
	}
	return dto, nil
}
