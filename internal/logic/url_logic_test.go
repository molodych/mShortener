package logic

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"testing"
)

type repository struct{}

func (*repository) GetByKey(ctx context.Context, key string) (result bson.M, err error) {
	if key == "12345678" {
		return bson.M{"_id": "123456789", "key": "12345678", "url": "http://w3.org"}, nil
	}
	return nil, errors.New("We have a problem")
}
func (*repository) Create(ctx context.Context, url map[string]string) (err error) {
	fmt.Println(url["key"])
	fmt.Println(url["url"])
	if url["key"] == "66c64974" && url["url"] == "http://w3.org" {
		return nil
	}
	return errors.New("We have a problem")
}

func TestLogic(t *testing.T) {
	logic := NewUrlLogic(&repository{})
	t.Run("TestGetByKey positive", func(t *testing.T) {
		_, err := logic.GetByKey(context.Background(), "12345678")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("TestGetByKey negative", func(t *testing.T) {
		_, err := logic.GetByKey(context.Background(), "87654321")
		if err == nil {
			t.Error("Different value!")
		}
	})
	t.Run("Create positive", func(t *testing.T) {
		_, err := logic.Create(context.Background(), "http://w3.org")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Create positive (empty http case)", func(t *testing.T) {
		_, err := logic.Create(context.Background(), "w3.org")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Create negative", func(t *testing.T) {
		_, err := logic.Create(context.Background(), "http://w3")
		if err == nil {
			t.Error("Different value!")
		}
	})
}
