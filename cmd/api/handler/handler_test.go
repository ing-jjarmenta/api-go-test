package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	domain "github.com/ing-jjarmenta/api-go-test/internal/domain/task"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/jsonencodec"
	service "github.com/ing-jjarmenta/api-go-test/internal/service/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetAllTasks(t *testing.T) {
	expectedTasks := []domain.Task{
		{
			ID:          bson.NewObjectID(),
			Title:       "Revisión de contrato",
			Description: "Analizar y validar las cláusulas del contrato con el cliente.",
			Status:      "pending",
			AssignedTo:  "Laura Pérez",
			DueDate:     "2025-08-15",
		},
	}
	tests := []struct {
		name    string
		request func() *http.Request
		mocks   func(*httptest.ResponseRecorder) (*service.MockTaskService, *jsonencodec.MockEncoder)
		asserts func(*httptest.ResponseRecorder)
	}{
		{
			name: "success",
			request: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
			},
			mocks: func(w *httptest.ResponseRecorder) (*service.MockTaskService, *jsonencodec.MockEncoder) {
				service := new(service.MockTaskService)
				service.On("GetAll", mock.Anything).Return(expectedTasks, nil)

				jsonEncoder := new(jsonencodec.MockEncoder)
				jsonEncoder.On("Encode", expectedTasks).Return(nil).Run(func(args mock.Arguments) {
					data, _ := json.Marshal(args.Get(0))
					w.Write(data)
				})

				return service, jsonEncoder
			},
			asserts: func(w *httptest.ResponseRecorder) {
				var result []domain.Task
				err := json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, expectedTasks, result)
			},
		},
		{
			name: "failed service get all tasks",
			request: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
			},
			mocks: func(w *httptest.ResponseRecorder) (*service.MockTaskService, *jsonencodec.MockEncoder) {
				service := new(service.MockTaskService)
				service.On("GetAll", mock.Anything).Return([]domain.Task{}, errors.New("service error get tasks"))

				return service, nil
			},
			asserts: func(w *httptest.ResponseRecorder) {
				result := w.Result()
				defer result.Body.Close()

				assert.Equal(t, http.StatusInternalServerError, w.Code)

				body, err := io.ReadAll(result.Body)
				assert.NoError(t, err)

				assert.Contains(t, string(body), "failed to get tasks: service error get tasks")
			},
		},
		{
			name: "failed json encode",
			request: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
			},
			mocks: func(w *httptest.ResponseRecorder) (*service.MockTaskService, *jsonencodec.MockEncoder) {
				service := new(service.MockTaskService)
				service.On("GetAll", mock.Anything).Return(expectedTasks, nil)

				jsonEncoder := new(jsonencodec.MockEncoder)
				jsonEncoder.On("Encode", expectedTasks).Return(errors.New("encode error"))

				return service, jsonEncoder
			},
			asserts: func(w *httptest.ResponseRecorder) {
				result := w.Result()
				defer result.Body.Close()

				assert.Equal(t, http.StatusInternalServerError, w.Code)

				body, err := io.ReadAll(result.Body)
				assert.NoError(t, err)

				assert.Contains(t, string(body), "failed to encode response: encode error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			service, jsonEncoder := tt.mocks(w)
			encoderFactory := func(w io.Writer) handler.Encoder {
				return jsonEncoder
			}

			taskHandler := newTaskHandler(t, service, encoderFactory)
			taskHandler.GetAllTasks(w, tt.request())
			tt.asserts(w)
			service.AssertCalled(t, "GetAll", mock.Anything)
			if jsonEncoder != nil {
				jsonEncoder.AssertCalled(t, "Encode", expectedTasks)
			}
		})
	}
}

func newTaskHandler(t *testing.T, mockTaskService *service.MockTaskService, encoderFactory handler.EncoderFactory) handler.TaskHandler {
	taskHandler := handler.NewTaskHandler(mockTaskService, encoderFactory)
	assert.NotNil(t, taskHandler)

	return *taskHandler
}
