package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wgir/gapsi-todo/internal/application"
	"github.com/wgir/gapsi-todo/internal/domain"
	"go.uber.org/zap"
)

type TaskHandler struct {
	service application.TaskService
	logger  *zap.Logger
}

func NewTaskHandler(service application.TaskService, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{service, logger}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTask(r.Context(), &task); err != nil {
		h.logger.Error("failed to create task", zap.Error(err))
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
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.logger.Error("failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateTask(r.Context(), id, &task); err != nil {
		h.logger.Error("failed to update task", zap.String("id", id), zap.Error(err))
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
