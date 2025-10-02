package routes_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	"github.com/ing-jjarmenta/api-go-test/cmd/api/routes"
	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoutesAPIV1(t *testing.T) {
	jsonExpected := `[{"id":"123","title":"fake task"}]`
	mockTask := new(handler.MockTaskHandler)

	mockTask.On("GetAllTasks", mock.Anything, mock.Anything).
		Return(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(jsonExpected))
		})

	handlers := dependencies.Handlers{Task: mockTask}
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, handlers)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/v1/tasks")
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.JSONEq(t, jsonExpected, string(body))

	mockTask.AssertCalled(t, "GetAllTasks", mock.Anything, mock.Anything)
}

func TestPing(t *testing.T) {
	handlers := dependencies.Handlers{Task: new(handler.MockTaskHandler)}
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, handlers)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/ping")
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "pong", string(body))
}
