package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UrlLogic interface {
	GetByKey(ctx context.Context, key string) (result map[string]interface{}, err error)
	Create(ctx context.Context, url string) (result string, err error)
}
type UrlRepository interface {
	GetByKey(ctx context.Context, key string) (result bson.M, err error)
	Create(ctx context.Context, url map[string]string) (err error)
}
