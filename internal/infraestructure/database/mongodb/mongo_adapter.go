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

func (ac *AdapterClient) Database(name string, opts ...options.Lister[options.DatabaseOptions]) MongoDatabase {
	return &AdapterDatabase{ac.Client.Database(name, opts...)}
}

func (ac *AdapterClient) Disconnect(ctx context.Context) error {
	return ac.Client.Disconnect(ctx)
}

func (ad *AdapterDatabase) Collection(name string, opts ...options.Lister[options.CollectionOptions]) MongoCollection {
	return &AdapterCollection{ad.Database.Collection(name, opts...)}
}

func (ac *AdapterCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	return ac.Collection.Find(ctx, filter, opts...)
}
