package task

import (
	"context"
	"errors"
	"testing"

	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name    string
		mocks   func(ctx context.Context) (*mongodb.MockCollection, *mongodb.MockCursor)
		asserts func([]domain.Task, error)
	}{
		{
			name: "find error",
			mocks: func(ctx context.Context) (*mongodb.MockCollection, *mongodb.MockCursor) {
				collection := new(mongodb.MockCollection)
				collection.
					On("Find", ctx, mock.Anything).
					Return((*mongodb.MockCursor)(nil), errors.New("find error collection"))

				return collection, nil
			},
			asserts: func(task []domain.Task, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "find error collection")
				assert.Nil(t, task)
			},
		},
		{
			name: "decode error",
			mocks: func(ctx context.Context) (*mongodb.MockCollection, *mongodb.MockCursor) {
				cursor := new(mongodb.MockCursor)
				cursor.On("Next", ctx).Return(true).Once()
				cursor.On("Decode", mock.Anything).
					Return(domain.Task{}, errors.New("decode error")).
					Once()

				collection := new(mongodb.MockCollection)
				collection.
					On("Find", ctx, mock.Anything).
					Return(cursor, nil)

				return collection, cursor
			},
			asserts: func(task []domain.Task, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "decode error")
				assert.Nil(t, task)
			},
		},
		{
			name: "cursor error",
			mocks: func(ctx context.Context) (*mongodb.MockCollection, *mongodb.MockCursor) {
				cursor := new(mongodb.MockCursor)
				cursor.On("Next", ctx).Return(false).Once()
				cursor.On("Err").Return(errors.New("cursor error")).Once()

				collection := new(mongodb.MockCollection)
				collection.
					On("Find", ctx, mock.Anything).
					Return(cursor, nil)

				return collection, cursor
			},
			asserts: func(task []domain.Task, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "cursor error")
				assert.Nil(t, task)
			},
		},
		{
			name: "success",
			mocks: func(ctx context.Context) (*mongodb.MockCollection, *mongodb.MockCursor) {
				cursor := new(mongodb.MockCursor)
				cursor.On("Next", ctx).Return(true).Once()
				cursor.On("Decode", mock.Anything).
					Return(domain.Task{Title: "Tarea 1"}, nil).
					Once()
				cursor.On("Next", ctx).Return(false).Once()
				cursor.On("Err").Return(nil).Once()

				collection := new(mongodb.MockCollection)
				collection.
					On("Find", ctx, mock.Anything).
					Return(cursor, nil)

				return collection, cursor
			},
			asserts: func(task []domain.Task, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, task)
				assert.Len(t, task, 1)
				assert.Equal(t, "Tarea 1", task[0].Title)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			mockCollection, mockCursor := tt.mocks(ctx)
			repository := newTaskRepository(t, mockCollection)
			tt.asserts(repository.GetAll(t.Context()))
			mockCollection.AssertExpectations(t)
			if mockCursor != nil {
				mockCursor.AssertExpectations(t)
			}
		})
	}
}

func newTaskRepository(t *testing.T, mockCollection *mongodb.MockCollection) TaskRepository {
	taskRepository := NewTaskRepository(mockCollection)
	assert.NotNil(t, taskRepository)

	return *taskRepository
}
