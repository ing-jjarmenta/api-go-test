package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/jsonencodec"
	repository "github.com/ing-jjarmenta/api-go-test/internal/repository/task"
	service "github.com/ing-jjarmenta/api-go-test/internal/service/task"
)

func main() {
	ctx := context.Background()

	client, err := mongodb.NewMongoClient(ctx)
	if err != nil {
		log.Fatal("error conectando a mongo:", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println("Error en la desconexión")
			log.Fatal(err)
		}
	}()

	tasksCollection := mongodb.TasksCollection(client)
	repository := repository.NewTaskRepository(tasksCollection)
	service := service.NewTaskService(repository)
	taskHandler := handler.NewTaskHandler(service, jsonencodec.NewJSONEncoderFactory())
	//log.Println(service.GetAll(ctx))

	http.HandleFunc("/api/tasks", taskHandler.GetAllTasks)
	http.HandleFunc("/api/hello", helloHandler)

	log.Println("Server 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("Error desde el serve")
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hola desde la API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
