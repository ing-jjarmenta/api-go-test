package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoClient(ctx context.Context) (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Println("Error en la conexión")
		log.Fatal(err)

		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Println("Error en el Ping")
		log.Fatal(err)

		return nil, err
	}

	log.Println("Conectado a MongoDB")

	return client, nil
}

func TasksCollection(client *mongo.Client) *mongo.Collection {
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "apidb"
	}

	return client.Database(dbName).Collection("tasks")
}
