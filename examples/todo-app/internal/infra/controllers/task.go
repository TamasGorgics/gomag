package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/TamasGorgics/gomag/examples/todo-app/internal/app"
	"github.com/TamasGorgics/gomag/examples/todo-app/internal/domain"
	"github.com/TamasGorgics/gomag/tx"
)

type TaskController struct {
	taskService *app.TaskService
	db          *sql.DB
}

func NewTaskController(taskService *app.TaskService, db *sql.DB) *TaskController {
	return &TaskController{taskService: taskService, db: db}
}

func (c *TaskController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.listTasks(w, r)
	case http.MethodPost:
		c.createTask(w, r)
	case http.MethodPut:
		c.updateTask(w, r)
	case http.MethodDelete:
		c.deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func (c *TaskController) listTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []*domain.Task

	err := tx.Exec(
		r.Context(),
		c.db,
		sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  true,
		},
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			tasks, err = c.taskService.ListTasks(ctx, tx)
			return err
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ActorName   string `json:"actor_name"`
}

func (c *TaskController) createTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if req.Description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}
	if req.ActorName == "" {
		http.Error(w, "Actor name is required", http.StatusBadRequest)
		return
	}

	var task *domain.Task

	err := tx.Exec(
		r.Context(),
		c.db,
		sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false},
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			task, err = c.taskService.CreateTask(ctx, tx, req.Title, req.Description, req.ActorName)
			return err
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

type updateTaskRequest struct {
	Completed bool `json:"completed"`
}

func (c *TaskController) updateTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	if idParam == "" {
		http.Error(w, "ID path parameter is required", http.StatusBadRequest)
		return
	}
	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = tx.Exec(
		r.Context(),
		c.db,
		sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		},
		func(ctx context.Context, tx *sql.Tx) error {
			if req.Completed {
				return c.taskService.CompleteTask(ctx, tx, parsedID)
			}
			return c.taskService.IncompleteTask(ctx, tx, parsedID)
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *TaskController) deleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	if idParam == "" {
		http.Error(w, "ID path parameter is required", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = tx.Exec(
		r.Context(),
		c.db,
		sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		},
		func(ctx context.Context, tx *sql.Tx) error {
			return c.taskService.DeleteTask(ctx, tx, parsedID)
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
