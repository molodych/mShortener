package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"mShorter/pkg/common"
	"testing"
)

func TestRepository(t *testing.T) {
	repository := NewUrlRepository()
	collection := common.GetMongoClient("mShorterDB-test")
	ctx := context.WithValue(context.Background(), "mongo-client", collection)

	defer collection.Collection("urls").DeleteMany(ctx, bson.M{"key": "66c64974"})
	t.Run("Create positive", func(t *testing.T) {
		err := repository.Create(ctx, map[string]string{"key": "66c64974", "url": "http://w3.org"})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("GetByKey positive", func(t *testing.T) {
		result, err := repository.GetByKey(ctx, "66c64974")
		if err != nil {
			t.Error(err)
		}
		if result["key"] != "66c64974" || result["url"] != "http://w3.org" {
			t.Error("Different value!")
		}
	})
}
