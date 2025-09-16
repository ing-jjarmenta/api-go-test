package mongodb

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestTasksCollection(t *testing.T) {
	tests := []struct {
		name       string
		envDB      string
		expectedDB string
	}{
		{
			name:       "default DB",
			envDB:      "",
			expectedDB: "apidb",
		},
		{
			name:       "custom DB",
			envDB:      "customdb",
			expectedDB: "customdb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envDB == "" {
				os.Unsetenv("MONGO_DB")
			} else {
				os.Setenv("MONGO_DB", tt.envDB)
			}

			mockCollection := new(MockCollection)
			mockDataBase := new(MockDataBase)
			mockClient := new(MockClient)

			mockClient.On("Database", tt.expectedDB).Return(mockDataBase)
			mockDataBase.On("Collection", "tasks").Return(mockCollection)

			collection := TasksCollection(mockClient)

			assert.Equal(t, mockCollection, collection)
			mockClient.AssertExpectations(t)
			mockDataBase.AssertExpectations(t)
		})
	}
}

func TestNewMongoClient(t *testing.T) {
	tests := []struct {
		name        string
		mockConnect func(opts ...*options.ClientOptions) (*mongo.Client, error)
		mockPing    func(ctx context.Context, client *mongo.Client) error
		asserts     func(client MongoClient, err error)
	}{
		{
			name: "connect error",
			mockConnect: func(opts ...*options.ClientOptions) (*mongo.Client, error) {
				return nil, errors.New("mock connect error")
			},
			mockPing: pingFunc,
			asserts: func(client MongoClient, err error) {
				assert.Nil(t, client)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "mock connect error")
			},
		},
		{
			name: "ping error",
			mockConnect: func(opts ...*options.ClientOptions) (*mongo.Client, error) {
				return &mongo.Client{}, nil
			},
			mockPing: func(ctx context.Context, client *mongo.Client) error {
				return errors.New("mock ping error")
			},
			asserts: func(client MongoClient, err error) {
				assert.Nil(t, client)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "mock ping error")
			},
		},
		{
			name: "success",
			mockConnect: func(opts ...*options.ClientOptions) (*mongo.Client, error) {
				return &mongo.Client{}, nil
			},
			mockPing: func(ctx context.Context, client *mongo.Client) error {
				return nil
			},
			asserts: func(client MongoClient, err error) {
				assert.NotNil(t, client)
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// backup originales
			origConnect := connectFunc
			origPing := pingFunc
			defer func() {
				connectFunc = origConnect
				pingFunc = origPing
			}()

			connectFunc = tt.mockConnect
			pingFunc = tt.mockPing

			tt.asserts(NewMongoClient(context.Background()))
		})
	}
}

func TestWrappers(t *testing.T) {
	// No validará conexión real, solo que las funciones existen y son invocables.
	_, _ = connectFunc(options.Client()) // ignoro error
	_ = pingFunc(context.Background(), &mongo.Client{})
}
