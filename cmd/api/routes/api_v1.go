package routes

import (
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
)

func RoutesAPIV1(handlers dependencies.Handlers) http.Handler {
	apiV1 := http.NewServeMux()

	apiV1.HandleFunc("/tasks", handlers.Task.GetAllTasks)

	return apiV1
}
