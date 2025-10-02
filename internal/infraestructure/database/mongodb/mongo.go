package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// Patrón adapter para desacoplar Mongo y poder realizar pruebas unitarias independientes

type MongoClient interface {
	Database(name string, opts ...options.Lister[options.DatabaseOptions]) MongoDatabase
	Disconnect(ctx context.Context) error
}

type MongoDatabase interface {
	Collection(name string, opts ...options.Lister[options.CollectionOptions]) MongoCollection
}

type MongoCollection interface {
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (MongoCursor, error)
}

type MongoCursor interface {
	Next(ctx context.Context) bool
	Decode(val any) error
	Err() error
	Close(ctx context.Context) error
}

// Wrappers necesarios para simular y facilitar el comportamiento en pruebas unitarias

var ConnectFunc = func(opts ...*options.ClientOptions) (*mongo.Client, error) {
	return mongo.Connect(opts...)
}

var PingFunc = func(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, readpref.Primary())
}

func NewMongoClient(ctx context.Context) (MongoClient, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := ConnectFunc(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error conectando a MongoDB: %w", err)
	}

	if err = PingFunc(ctx, client); err != nil {
		return nil, fmt.Errorf("error haciendo ping a MongoDB: %w", err)
	}

	return &AdapterClient{client}, nil
}

func TasksCollection(client MongoClient) MongoCollection {
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "apidb"
	}

	return client.Database(dbName).Collection("tasks")
}
