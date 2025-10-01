package server

import (
	"context"
	"log"
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
)

func Run() error {
	ctx := context.Background()
	client, err := dependencies.ResolveMongoClient(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mux := http.NewServeMux()
	RegisterRoutes(mux, dependencies.ResolveHandlers(client))

	log.Println("Server running on :8081")

	return http.ListenAndServe(":8081", mux)
}
