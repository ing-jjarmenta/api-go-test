package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	mongoDBConection()
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

func mongoDBConection() {
	ctx := context.Background()

	client, err := mongodb.NewMongoClient(ctx)
	if err != nil {
		log.Fatal("error conectando a mongo:", err)
	}

	tasksCollection := mongodb.TasksCollection(client)
	cursor, err := tasksCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("Error Find")
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		log.Println("task")
		log.Println(result)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error en el cursor")
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println("Error en la desconexión")
			log.Fatal(err)
		}
	}()
}
