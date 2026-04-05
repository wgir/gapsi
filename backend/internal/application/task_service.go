package application

import (
	"context"
	"errors"

	"github.com/wgir/gapsi-todo/internal/domain"
)

// TaskService handles business logic related to tasks.
type TaskService interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context, status domain.TaskStatus, limit int, lastID string) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	UpdateTask(ctx context.Context, id string, task *domain.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type taskService struct {
	repo domain.TaskRepository
}

// NewTaskService creates a new implementation of TaskService.
func NewTaskService(repo domain.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) CreateTask(ctx context.Context, task *domain.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.Status == "" {
		task.Status = domain.StatusTodo
	}
	return s.repo.Create(ctx, task)
}

func (s *taskService) GetAllTasks(ctx context.Context, status domain.TaskStatus, limit int, lastID string) ([]domain.Task, error) {
	if limit <= 0 {
		limit = 10
	}
	return s.repo.GetAll(ctx, status, limit, lastID)
}

func (s *taskService) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *taskService) UpdateTask(ctx context.Context, id string, task *domain.Task) error {
	existingTask, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Update fields
	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Description != "" {
		existingTask.Description = task.Description
	}
	if task.Status != "" {
		existingTask.Status = task.Status
	}

	return s.repo.Update(ctx, existingTask)
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
