package dependencies

import (
	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/jsonencodec"
	repository "github.com/ing-jjarmenta/api-go-test/internal/repository/task"
	service "github.com/ing-jjarmenta/api-go-test/internal/service/task"
)

func resolveTaskCollection(client mongodb.MongoClient) mongodb.MongoCollection {
	return mongodb.TasksCollection(client)
}

func resolveTaskRepository(client mongodb.MongoClient) *repository.TaskRepository {
	return repository.NewTaskRepository(resolveTaskCollection(client))
}

func resolveTaskService(client mongodb.MongoClient) *service.TaskService {
	return service.NewTaskService(resolveTaskRepository(client))
}

func ResolveTaskHandler(client mongodb.MongoClient) *handler.TaskHandler {
	return handler.NewTaskHandler(resolveTaskService(client), jsonencodec.NewJSONEncoderFactory())
}
