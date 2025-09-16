package mongodb

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MockClient struct {
	mock.Mock
}

type MockDataBase struct {
	mock.Mock
}

type MockCollection struct {
	mock.Mock
}

func (mc *MockClient) Database(name string, _ ...options.Lister[options.DatabaseOptions]) MongoDatabase {
	args := mc.Called(name)

	return args.Get(0).(MongoDatabase)
}

func (mc *MockClient) Disconnect(ctx context.Context) error {
	args := mc.Called(ctx)

	return args.Error(0)
}

func (mdb *MockDataBase) Collection(name string, _ ...options.Lister[options.CollectionOptions]) MongoCollection {
	args := mdb.Called(name)

	return args.Get(0).(MongoCollection)
}

func (mcoll *MockCollection) Find(ctx context.Context, filter any, _ ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	args := mcoll.Called(ctx, filter)

	return args.Get(0).(*mongo.Cursor), args.Error(1)
}
