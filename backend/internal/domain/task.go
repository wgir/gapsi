package domain

import (
	"context"
	"time"
)

// TaskStatus represents the current state of a task.
type TaskStatus string

const (
	StatusTodo      TaskStatus = "TODO"
	StatusDone      TaskStatus = "DONE"
	StatusCancelled TaskStatus = "CANCELLED"
)

// Task represents a todo item in the system.
type Task struct {
	ID          string     `json:"id" firestore:"-"` // ID is not stored in the document body in Firestore
	Title       string     `json:"title" firestore:"title"`
	Description string     `json:"description" firestore:"description"`
	Status      TaskStatus `json:"status" firestore:"status"`
	CreatedAt   time.Time  `json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" firestore:"updated_at"`
}

// TaskRepository defines the interface for interacting with task storage.
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetAll(ctx context.Context, status TaskStatus) ([]Task, error)
	GetByID(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
}
