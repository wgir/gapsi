package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/wgir/gapsi-todo/internal/application"
	"github.com/wgir/gapsi-todo/internal/domain"
	"go.uber.org/zap"
)

type createTaskResult struct {
	statusCode    int
	body          string
	serviceCalled bool
	receivedTask  domain.Task
}

type createTaskChecker func(value createTaskResult, err error)

type mockTaskService struct {
	createTaskFn func(ctx context.Context, task *domain.Task) error
	updateTaskFn func(ctx context.Context, id string, task *domain.Task) error

	createTaskCalled bool
	receivedTask     domain.Task
	updateTaskCalled bool
	updateTaskID     string
	updateTask       domain.Task
}

func (m *mockTaskService) CreateTask(ctx context.Context, task *domain.Task) error {
	m.createTaskCalled = true
	m.receivedTask = *task

	if m.createTaskFn != nil {
		return m.createTaskFn(ctx, task)
	}

	return nil
}

func (m *mockTaskService) GetAllTasks(ctx context.Context, status domain.TaskStatus, limit int, lastID string) ([]domain.Task, error) {
	return nil, nil
}

func (m *mockTaskService) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	return nil, nil
}

func (m *mockTaskService) UpdateTask(ctx context.Context, id string, task *domain.Task) error {
	m.updateTaskCalled = true
	m.updateTaskID = id
	m.updateTask = *task

	if m.updateTaskFn != nil {
		return m.updateTaskFn(ctx, id, task)
	}

	return nil
}

func (m *mockTaskService) DeleteTask(ctx context.Context, id string) error {
	return nil
}

func executeCreateTask(handler *TaskHandler, service *mockTaskService, requestBody string) (createTaskResult, error) {
	req := httptest.NewRequest(http.MethodPost, "/tasks/", strings.NewReader(requestBody))
	rec := httptest.NewRecorder()

	handler.CreateTask(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return createTaskResult{}, err
	}

	return createTaskResult{
		statusCode:    resp.StatusCode,
		body:          strings.TrimSpace(string(bodyBytes)),
		serviceCalled: service.createTaskCalled,
		receivedTask:  service.receivedTask,
	}, nil
}

type updateTaskResult struct {
	statusCode      int
	body            string
	serviceCalled   bool
	receivedTaskID  string
	receivedTask    domain.Task
}

type updateTaskChecker func(value updateTaskResult, err error)

func executeUpdateTask(handler *TaskHandler, service *mockTaskService, taskID string, requestBody string) (updateTaskResult, error) {
	req := httptest.NewRequest(http.MethodPut, "/tasks/"+taskID, strings.NewReader(requestBody))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, &chi.Context{
		URLParams: chi.RouteParams{
			Keys:   []string{"id"},
			Values: []string{taskID},
		},
	}))
	rec := httptest.NewRecorder()

	handler.UpdateTask(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return updateTaskResult{}, err
	}

	return updateTaskResult{
		statusCode:     resp.StatusCode,
		body:           strings.TrimSpace(string(bodyBytes)),
		serviceCalled:  service.updateTaskCalled,
		receivedTaskID: service.updateTaskID,
		receivedTask:   service.updateTask,
	}, nil
}

func TestTaskHandler_CreateTask_WhenThen_TableDriven(t *testing.T) {
	testCases := []struct {
		name        string
		requestBody string
		serviceFn   func(ctx context.Context, task *domain.Task) error
		checker     createTaskChecker
	}{
		{
			name:        "When request body has invalid JSON Then it returns bad request and does not call service",
			requestBody: `{"title":`,
			checker: func(value createTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.True(
					t,
					strings.Contains(value.body, "invalid character") || strings.Contains(value.body, "unexpected EOF"),
				)
				assert.False(t, value.serviceCalled)
			},
		},
		{
			name:        "When title is empty Then it returns bad request",
			requestBody: `{"title":"","description":"valid description"}`,
			checker: func(value createTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, application.ErrTitleRequired.Error())
			},
		},
		{
			name:        "When description is empty Then it returns bad request",
			requestBody: `{"title":"valid title","description":""}`,
			checker: func(value createTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, application.ErrDescriptionRequired.Error())
			},
		},
		{
			name:        "When service returns an error Then it returns internal server error",
			requestBody: `{"title":"Task with service failure","description":"desc"}`,
			serviceFn: func(ctx context.Context, task *domain.Task) error {
				return errors.New("create failed")
			},
			checker: func(value createTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, value.statusCode)
				assert.True(t, value.serviceCalled)
				assert.Contains(t, value.body, "create failed")
				assert.Equal(t, "Task with service failure", value.receivedTask.Title)
			},
		},
		{
			name:        "When request body is valid Then it returns created task",
			requestBody: `{"title":"My new task","description":"clean architecture"}`,
			serviceFn: func(ctx context.Context, task *domain.Task) error {
				task.ID = "task-1"
				task.Status = domain.StatusTodo
				return nil
			},
			checker: func(value createTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusCreated, value.statusCode)
				assert.True(t, value.serviceCalled)
				assert.Equal(t, "My new task", value.receivedTask.Title)

				var responseTask domain.Task
				unmarshalErr := json.Unmarshal([]byte(value.body), &responseTask)
				assert.NoError(t, unmarshalErr)
				assert.Equal(t, "task-1", responseTask.ID)
				assert.Equal(t, "My new task", responseTask.Title)
				assert.Equal(t, "clean architecture", responseTask.Description)
				assert.Equal(t, domain.StatusTodo, responseTask.Status)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			service := &mockTaskService{createTaskFn: tc.serviceFn}
			handler := NewTaskHandler(service, zap.NewNop())

			value, err := executeCreateTask(handler, service, tc.requestBody)
			tc.checker(value, err)
		})
	}
}

func TestTaskHandler_UpdateTask_WhenThen_TableDriven(t *testing.T) {
	testCases := []struct {
		name        string
		taskID      string
		requestBody string
		serviceFn   func(ctx context.Context, id string, task *domain.Task) error
		checker     updateTaskChecker
	}{
		{
			name:        "When request body has invalid JSON Then it returns bad request and does not call service",
			taskID:      "task-1",
			requestBody: `{"title":`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.True(
					t,
					strings.Contains(value.body, "invalid character") || strings.Contains(value.body, "unexpected EOF"),
				)
				assert.False(t, value.serviceCalled)
			},
		},
		{
			name:        "When title is empty Then it returns bad request and does not call service",
			taskID:      "task-2",
			requestBody: `{"title":"   ","status":"TODO"}`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, application.ErrTitleRequired.Error())
			},
		},
		{
			name:        "When description is empty Then it returns bad request and does not call service",
			taskID:      "task-3",
			requestBody: `{"description":"   ","status":"DONE"}`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, application.ErrDescriptionRequired.Error())
			},
		},
		{
			name:        "When status is missing Then it returns bad request and does not call service",
			taskID:      "task-4",
			requestBody: `{"title":"Refactor tests"}`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, "status is required")
			},
		},
		{
			name:        "When status is invalid Then it returns bad request and does not call service",
			taskID:      "task-5",
			requestBody: `{"status":"INVALID"}`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, value.statusCode)
				assert.False(t, value.serviceCalled)
				assert.Contains(t, value.body, "status must be one of TODO, DONE, CANCELLED")
			},
		},
		{
			name:        "When service returns an error Then it returns internal server error",
			taskID:      "task-6",
			requestBody: `{"title":"Refactor tests","status":"DONE"}`,
			serviceFn: func(ctx context.Context, id string, task *domain.Task) error {
				return errors.New("update failed")
			},
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusInternalServerError, value.statusCode)
				assert.True(t, value.serviceCalled)
				assert.Contains(t, value.body, "update failed")
				assert.Equal(t, "task-6", value.receivedTaskID)
				assert.Equal(t, "Refactor tests", value.receivedTask.Title)
			},
		},
		{
			name:        "When request body is valid Then it returns ok and calls service with normalized fields",
			taskID:      "task-7",
			requestBody: `{"title":"  Updated Title  ","description":" Updated Description ","status":"DONE"}`,
			checker: func(value updateTaskResult, err error) {
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, value.statusCode)
				assert.Equal(t, "", value.body)
				assert.True(t, value.serviceCalled)
				assert.Equal(t, "task-7", value.receivedTaskID)
				assert.Equal(t, "Updated Title", value.receivedTask.Title)
				assert.Equal(t, "Updated Description", value.receivedTask.Description)
				assert.Equal(t, domain.StatusDone, value.receivedTask.Status)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			service := &mockTaskService{updateTaskFn: tc.serviceFn}
			handler := NewTaskHandler(service, zap.NewNop())

			value, err := executeUpdateTask(handler, service, tc.taskID, tc.requestBody)
			tc.checker(value, err)
		})
	}
}
