package server

import (
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
)

func RegisterRoutes(mux *http.ServeMux, handlers dependencies.Handlers) {
	mux.HandleFunc("/api/tasks", handlers.Task.GetAllTasks)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`pong`))
	})
}
