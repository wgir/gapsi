package application

import (
	"context"
	"errors"
	"testing"

	"github.com/wgir/gapsi-todo/internal/domain"
)

// MockRepository is a simple mock for domain.TaskRepository.
type MockRepository struct {
	tasks map[string]domain.Task
}

func (m *MockRepository) Create(ctx context.Context, task *domain.Task) error {
	m.tasks[task.ID] = *task
	return nil
}

func (m *MockRepository) GetAll(ctx context.Context, status domain.TaskStatus) ([]domain.Task, error) {
	var results []domain.Task
	for _, t := range m.tasks {
		if status == "" || t.Status == status {
			results = append(results, t)
		}
	}
	return results, nil
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	t, ok := m.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &t, nil
}

func (m *MockRepository) Update(ctx context.Context, task *domain.Task) error {
	m.tasks[task.ID] = *task
	return nil
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	delete(m.tasks, id)
	return nil
}

func TestTaskService_CreateTask(t *testing.T) {
	repo := &MockRepository{tasks: make(map[string]domain.Task)}
	service := NewTaskService(repo)

	task := &domain.Task{Title: "Test Task"}
	err := service.CreateTask(context.Background(), task)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if task.Status != domain.StatusTodo {
		t.Errorf("Expected status TODO, got %s", task.Status)
	}
}

func TestTaskService_CreateTask_NoTitle(t *testing.T) {
	repo := &MockRepository{tasks: make(map[string]domain.Task)}
	service := NewTaskService(repo)

	task := &domain.Task{Title: ""}
	err := service.CreateTask(context.Background(), task)

	if err == nil {
		t.Error("Expected error for empty title, got nil")
	}
}
