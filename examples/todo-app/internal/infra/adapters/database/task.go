package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/TamasGorgics/gomag/examples/todo-app/internal/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, tx *sql.Tx, task *domain.Task) (*domain.Task, error) {
	_, err := tx.ExecContext(ctx, "INSERT INTO tasks (id, title, description, completed, created_at, updated_at, actor_name) VALUES (?, ?, ?, ?, ?, ?, ?)", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt, task.UpdatedAt, task.ActorName)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) Get(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*domain.Task, error) {
	row := tx.QueryRowContext(ctx, "SELECT id, title, description, completed, created_at, updated_at, actor_name, version FROM tasks WHERE id = ?", id)
	var task domain.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.ActorName, &task.Version)
	return &task, err
}

func (r *TaskRepository) Update(ctx context.Context, tx *sql.Tx, task *domain.Task) error {
	_, err := tx.ExecContext(ctx, "UPDATE tasks SET title = ?, description = ?, completed = ?, updated_at = ?, actor_name = ?, version = ? WHERE id = ? AND version = ?", task.Title, task.Description, task.Completed, task.UpdatedAt, task.ActorName, task.Version, task.ID, task.Version-1)
	return err
}

func (r *TaskRepository) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (r *TaskRepository) List(ctx context.Context, tx *sql.Tx) ([]*domain.Task, error) {
	rows, err := tx.QueryContext(ctx, "SELECT id, title, description, completed, created_at, updated_at, actor_name, version FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt, &task.ActorName, &task.Version)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
