package database

import (
	"context"
	"deathstar"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, uri, databaseName string) (Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	targetCollection := client.Database(databaseName).Collection("targets")
	return &dataStore{
		targetCollection: targetCollection,
	}, nil
}

type dataStore struct {
	targetCollection *mongo.Collection
}

func (d *dataStore) AddTargets(ctx context.Context, targets []deathstar.Target) error {
	_, err := d.targetCollection.InsertMany(ctx, convertToSliceInterface(targets))
	return err
}

func convertToSliceInterface(input []deathstar.Target) []interface{} {
	change := make([]interface{}, len(input))
	for i := range input {
		change[i] = input[i]
	}
	return change
}
