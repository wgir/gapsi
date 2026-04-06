package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/wgir/gapsi-todo/internal/application"
	"github.com/wgir/gapsi-todo/internal/domain"
	"go.uber.org/zap"
)

type TaskHandler struct {
	service application.TaskService
	logger  *zap.Logger
}

var taskValidator = validator.New()

var errInvalidTaskStatus = errors.New("status must be one of TODO, DONE, CANCELLED")
var errStatusRequired = errors.New("status is required")

type taskValidationOperation string

const (
	taskValidationCreate taskValidationOperation = "create"
	taskValidationUpdate taskValidationOperation = "update"
)

type taskRequest struct {
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Status      *domain.TaskStatus `json:"status"`
}

type createTaskValidation struct {
	Title       string            `validate:"required"`
	Description string            `validate:"required"`
	Status      domain.TaskStatus `validate:"omitempty,oneof=TODO DONE CANCELLED"`
}

type updateTaskValidation struct {
	Title       *string            `validate:"omitempty,min=1"`
	Description *string            `validate:"omitempty,min=1"`
	Status      *domain.TaskStatus `validate:"required,oneof=TODO DONE CANCELLED"`
}

func NewTaskHandler(service application.TaskService, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{service, logger}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var request taskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	normalizeTaskRequest(&request)
	if err := validateTaskRequest(request, taskValidationCreate); err != nil {
		h.logger.Error("invalid create task payload", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := request.toDomainTask()

	if err := h.service.CreateTask(r.Context(), &task); err != nil {
		h.logger.Error("failed to create task", zap.Error(err))
		if isTaskValidationError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	status := domain.TaskStatus(r.URL.Query().Get("status"))
	limitStr := r.URL.Query().Get("limit")
	lastID := r.URL.Query().Get("last_id")

	limit := 10
	if limitStr != "" {
		fmt.Sscanf(limitStr, "%d", &limit)
	}

	tasks, err := h.service.GetAllTasks(r.Context(), status, limit, lastID)
	if err != nil {
		h.logger.Error("failed to get all tasks", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := h.service.GetTaskByID(r.Context(), id)
	if err != nil {
		h.logger.Error("failed to get task by ID", zap.String("id", id), zap.Error(err))
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var request taskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	normalizeTaskRequest(&request)
	if err := validateTaskRequest(request, taskValidationUpdate); err != nil {
		h.logger.Error("invalid update task payload", zap.String("id", id), zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := request.toDomainTask()
	if err := h.service.UpdateTask(r.Context(), id, &task); err != nil {
		h.logger.Error("failed to update task", zap.String("id", id), zap.Error(err))
		if isTaskValidationError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		h.logger.Error("failed to delete task", zap.String("id", id), zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func normalizeTaskRequest(request *taskRequest) {
	if request.Title != nil {
		trimmedTitle := strings.TrimSpace(*request.Title)
		request.Title = &trimmedTitle
	}

	if request.Description != nil {
		trimmedDescription := strings.TrimSpace(*request.Description)
		request.Description = &trimmedDescription
	}
}

func (request taskRequest) toDomainTask() domain.Task {
	var task domain.Task

	if request.Title != nil {
		task.Title = *request.Title
	}
	if request.Description != nil {
		task.Description = *request.Description
	}
	if request.Status != nil {
		task.Status = *request.Status
	}

	return task
}

func validateTaskRequest(request taskRequest, operation taskValidationOperation) error {
	switch operation {
	case taskValidationCreate:
		createValidation := createTaskValidation{}
		if request.Title != nil {
			createValidation.Title = *request.Title
		}
		if request.Description != nil {
			createValidation.Description = *request.Description
		}
		if request.Status != nil {
			createValidation.Status = *request.Status
		}
		return mapTaskValidationErrors(taskValidator.Struct(createValidation))
	case taskValidationUpdate:
		updateValidation := updateTaskValidation{
			Title:       request.Title,
			Description: request.Description,
			Status:      request.Status,
		}
		return mapTaskValidationErrors(taskValidator.Struct(updateValidation))
	default:
		return nil
	}
}

func mapTaskValidationErrors(err error) error {
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "Title":
				return application.ErrTitleRequired
			case "Description":
				return application.ErrDescriptionRequired
			case "Status":
				if fieldError.Tag() == "required" {
					return errStatusRequired
				}
				return errInvalidTaskStatus
			}
		}
	}

	return err
}

func isTaskValidationError(err error) bool {
	return errors.Is(err, application.ErrTitleRequired) ||
		errors.Is(err, application.ErrDescriptionRequired) ||
		errors.Is(err, errStatusRequired) ||
		errors.Is(err, errInvalidTaskStatus)
}
