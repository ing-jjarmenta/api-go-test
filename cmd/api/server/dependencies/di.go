package dependencies

import (
	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
)

type Handlers struct {
	Task *handler.TaskHandler
}

func ResolveHandlers(client mongodb.MongoClient) Handlers {
	return Handlers{Task: ResolveTaskHandler(client)}
}
