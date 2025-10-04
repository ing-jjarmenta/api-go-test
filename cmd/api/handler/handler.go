package handler

import (
	"context"
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/dto/task/response"
	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]domain.Task, error)
}

type TaskHandler struct {
	service        TaskService
	encoderFactory EncoderFactory
}

func NewTaskHandler(service TaskService, encoderFactory EncoderFactory) *TaskHandler {
	return &TaskHandler{
		service:        service,
		encoderFactory: encoderFactory,
	}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "failed to get tasks: "+err.Error(), http.StatusInternalServerError)

		return
	}

	taskResponses := response.ToTaskResponses(tasks)
	if err := h.encoderFactory(w).Encode(taskResponses); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
