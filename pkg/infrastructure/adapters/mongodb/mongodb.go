package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    = context.TODO()
)

func Connect(connectionString string) error {
	clientOptions := options.Client().ApplyURI(connectionString)
	newClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("Error al conectar a MongoDB: %v", err)
	}

	err = newClient.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("Error al hacer ping a MongoDB: %v", err)
	}

	client = newClient
	return nil
}

func GetClient() *mongo.Client {
	return client
}

func Disconnect() error {
	if client != nil {
		return client.Disconnect(ctx)
	}
	return nil
}
