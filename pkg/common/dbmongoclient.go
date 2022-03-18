package common

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient(dbName string) *mongo.Database {
	const connectionString = "mongodb+srv://igmolodykh:Bujhmvjkjls_2410@cluster0.ixca0.mongodb.net/mShorterDB?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	return client.Database(dbName)
}
