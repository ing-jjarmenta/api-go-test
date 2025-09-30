package task

import (
	"context"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetAll(ctx context.Context) ([]domain.Task, error) {
	args := m.Called(ctx)

	return args.Get(0).([]domain.Task), args.Error(1)
}
