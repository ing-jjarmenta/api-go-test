package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Error en la conexión")
		log.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Error en el Ping")
		log.Fatal(err)
	}

	log.Println("Conectado a MongoDB")

	collection := client.Database("apidb").Collection("tasks")
	cursor, err := collection.Find(ctx, bson.D{})
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
