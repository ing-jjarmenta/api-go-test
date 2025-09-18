package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AdapterClient struct {
	*mongo.Client
}

type AdapterDatabase struct {
	*mongo.Database
}

type AdapterCollection struct {
	*mongo.Collection
}

type AdapterCursor struct {
	*mongo.Cursor
}

func (ac *AdapterClient) Database(name string, opts ...options.Lister[options.DatabaseOptions]) MongoDatabase {
	return &AdapterDatabase{ac.Client.Database(name, opts...)}
}

func (ac *AdapterClient) Disconnect(ctx context.Context) error {
	return ac.Client.Disconnect(ctx)
}

func (ad *AdapterDatabase) Collection(name string, opts ...options.Lister[options.CollectionOptions]) MongoCollection {
	return &AdapterCollection{ad.Database.Collection(name, opts...)}
}

func (ac *AdapterCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (MongoCursor, error) {
	cursor, err := ac.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return &AdapterCursor{Cursor: cursor}, nil
}

func (acursor *AdapterCursor) Next(ctx context.Context) bool {
	return acursor.Cursor.Next(ctx)
}

func (acursor *AdapterCursor) Decode(val any) error {
	return acursor.Cursor.Decode(val)
}

func (acursor *AdapterCursor) Err() error {
	return acursor.Cursor.Err()
}

func (acursor *AdapterCursor) Close(ctx context.Context) error {
	return acursor.Cursor.Close(ctx)
}
