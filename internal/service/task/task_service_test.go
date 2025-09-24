package task

import (
	"context"
	"testing"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	repository "github.com/ing-jjarmenta/api-go-test/internal/repository/task"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(ctx context.Context) *repository.MockTaskRepository
		asserts func([]domain.Task, error)
	}{
		{
			name: "success",
			mocks: func(ctx context.Context) *repository.MockTaskRepository {
				repository := new(repository.MockTaskRepository)
				repository.On("GetAll", ctx).Return([]domain.Task{}, nil).Once()

				return repository
			},
			asserts: func(tasks []domain.Task, err error) {
				assert.NoError(t, err)
				assert.Nil(t, err)
				assert.NotNil(t, tasks)
				assert.Len(t, tasks, 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			mockRepository := tt.mocks(ctx)
			service := newTaskService(t, mockRepository)
			tt.asserts(service.GetAll(ctx))
			mockRepository.AssertExpectations(t)
		})
	}
}

func newTaskService(t *testing.T, mockTaskRepository *repository.MockTaskRepository) TaskService {
	taskService := NewTaskService(mockTaskRepository)
	assert.NotNil(t, taskService)

	return *taskService
}
