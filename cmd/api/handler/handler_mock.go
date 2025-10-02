package handler

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockTaskHandler struct {
	mock.Mock
}

func (m *MockTaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	args := m.Called(w, r)
	if fn, ok := args.Get(0).(func(w http.ResponseWriter, r *http.Request)); ok {
		fn(w, r)
	}
}
