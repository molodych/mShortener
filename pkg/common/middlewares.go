package common

import (
	"context"
	"github.com/labstack/echo/v4"
)

func HTTPDBClientMongo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			ctx = context.WithValue(ctx, "mongo-client", GetMongoClient("mShorterDB"))
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
