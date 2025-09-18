package mongodb

import (
	"context"
	"fmt"
	"reflect"

	"github.com/stretchr/testify/mock"
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

type MockCursor struct {
	mock.Mock
}

// Mock Client
func (mc *MockClient) Database(name string, _ ...options.Lister[options.DatabaseOptions]) MongoDatabase {
	args := mc.Called(name)

	return args.Get(0).(MongoDatabase)
}

func (mc *MockClient) Disconnect(ctx context.Context) error {
	args := mc.Called(ctx)

	return args.Error(0)
}

// Mock DataBase
func (mdb *MockDataBase) Collection(name string, _ ...options.Lister[options.CollectionOptions]) MongoCollection {
	args := mdb.Called(name)

	return args.Get(0).(MongoCollection)
}

// Mock Collection
func (mcoll *MockCollection) Find(ctx context.Context, filter any, _ ...options.Lister[options.FindOptions]) (MongoCursor, error) {
	args := mcoll.Called(ctx, filter)

	return args.Get(0).(MongoCursor), args.Error(1)
}

// Mock Cursor
func (mcursor *MockCursor) Next(ctx context.Context) bool {
	args := mcursor.Called(ctx)

	return args.Bool(0)
}

func (mcursor *MockCursor) Decode(val any) error {
	args := mcursor.Called(val)
	v := args.Get(0)

	if v != nil {
		rv := reflect.ValueOf(val).Elem() // ej: domain.Task
		vv := reflect.ValueOf(v)

		switch {
		// Caso 1: valor directo (domain.Task → domain.Task)
		case vv.Type().AssignableTo(rv.Type()):
			rv.Set(vv)

		// Caso 2: puntero al valor (*domain.Task → domain.Task)
		case vv.Kind() == reflect.Pointer && vv.Elem().Type().AssignableTo(rv.Type()):
			rv.Set(vv.Elem())

		default:
			return fmt.Errorf("mockCursor.Decode: cannot assign %T to %T", v, rv.Interface())
		}
	}

	return args.Error(1)
}

func (mcursor *MockCursor) Err() error {
	args := mcursor.Called()

	return args.Error(0)
}

func (mcursor *MockCursor) Close(ctx context.Context) error {
	return nil
}
