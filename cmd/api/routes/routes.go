package routes

import (
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
)

func RegisterRoutes(mux *http.ServeMux, handlers dependencies.Handlers) {
	apiV1 := RoutesAPIV1(handlers)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`pong`))
	})
}
