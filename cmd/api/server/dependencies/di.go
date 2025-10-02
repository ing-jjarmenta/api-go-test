package dependencies

import (
	"context"
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
)

type TaskHandler interface {
	GetAllTasks(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Task TaskHandler
}

func ResolveMongoClient(ctx context.Context) (mongodb.MongoClient, error) {
	return mongodb.NewMongoClient(ctx)
}

func ResolveHandlers(client mongodb.MongoClient) Handlers {
	return Handlers{Task: ResolveTaskHandler(client)}
}
