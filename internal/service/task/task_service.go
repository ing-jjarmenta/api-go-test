package task

import (
	"context"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]domain.Task, error)
}

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) GetAll(ctx context.Context) ([]domain.Task, error) {
	return s.repository.GetAll(ctx)
}
