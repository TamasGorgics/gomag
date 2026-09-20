package app

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/TamasGorgics/gomag/examples/todo-app/internal/domain"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/infra/adapters/database"
)

type TaskService struct {
	taskRepository database.TaskRepository
}

func NewTaskService(taskRepository database.TaskRepository) *TaskService {
	return &TaskService{taskRepository: taskRepository}
}

func (s *TaskService) CreateTask(ctx context.Context, tx *sql.Tx, title, description, actorName string) (*domain.Task, error) {
	t := domain.NewTask(title, description, actorName)
	_, err := s.taskRepository.Create(ctx, tx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TaskService) CompleteTask(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	task, err := s.taskRepository.Get(ctx, tx, id)
	if err != nil {
		return err
	}
	task.MarkAsCompleted()
	return s.taskRepository.Update(ctx, tx, task)
}

func (s *TaskService) IncompleteTask(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	task, err := s.taskRepository.Get(ctx, tx, id)
	if err != nil {
		return err
	}
	task.MarkAsIncomplete()
	return s.taskRepository.Update(ctx, tx, task)
}

func (s *TaskService) ListTasks(ctx context.Context, tx *sql.Tx) ([]*domain.Task, error) {
	return s.taskRepository.List(ctx, tx)
}

func (s *TaskService) DeleteTask(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	return s.taskRepository.Delete(ctx, tx, id)
}
