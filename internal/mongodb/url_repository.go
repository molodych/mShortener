package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type urlRepository struct {
}

func NewUrlRepository() *urlRepository {
	return &urlRepository{}
}

func (*urlRepository) GetByKey(ctx context.Context, key string) (result bson.M, err error) {
	collection := ctx.Value("mongo-client").(*mongo.Client).Database("mShorterDB").Collection("urls")
	err = collection.FindOne(ctx, bson.M{"key": key}).Decode(&result)
	return
}

func (*urlRepository) Create(ctx context.Context, url map[string]string) (err error) {
	collection := ctx.Value("mongo-client").(*mongo.Client).Database("mShorterDB").Collection("urls")
	_, err = collection.InsertOne(ctx, bson.M{"key": url["key"], "url": url["url"]})
	return err
}
