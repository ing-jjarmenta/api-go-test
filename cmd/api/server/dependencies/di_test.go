package dependencies_test

import (
	"context"
	"os"
	"testing"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/handler"
	"github.com/ing-jjarmenta/api-go-test/cmd/api/server/dependencies"
	"github.com/ing-jjarmenta/api-go-test/internal/infraestructure/database/mongodb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestResolveMongoClient(t *testing.T) {
	origConnect := mongodb.ConnectFunc
	origPing := mongodb.PingFunc
	defer func() {
		mongodb.ConnectFunc = origConnect
		mongodb.PingFunc = origPing
	}()

	mongodb.ConnectFunc = func(opts ...*options.ClientOptions) (*mongo.Client, error) {
		return &mongo.Client{}, nil
	}
	mongodb.PingFunc = func(ctx context.Context, client *mongo.Client) error {
		return nil
	}

	client, err := dependencies.ResolveMongoClient(t.Context())

	assert.NoError(t, err)
	assert.NotNil(t, client)
	_, ok := client.(*mongodb.AdapterClient)
	assert.True(t, ok, "client debería ser *mongodb.AdapterClient")
}

func TestResolveHandlers(t *testing.T) {
	os.Setenv("MONGO_DB", "dbtest")
	mockMongoClient := new(mongodb.MockClient)
	mockDataBase := new(mongodb.MockDataBase)

	mockMongoClient.On("Database", "dbtest").Return(mockDataBase)
	mockDataBase.On("Collection", "tasks").Return(new(mongodb.MockCollection))

	handlers := dependencies.ResolveHandlers(mockMongoClient)

	assert.NotNil(t, handlers)
	assert.NotNil(t, handlers.Task)
	_, ok := handlers.Task.(*handler.TaskHandler)
	assert.True(t, ok, "handlers.Task debería ser *handler.TaskHandler")
}
