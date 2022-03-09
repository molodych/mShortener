package main

import (
	"github.com/labstack/echo/v4"
	"mShorter/internal/http"
	"mShorter/internal/logic"
	"mShorter/internal/mongodb"
	"mShorter/pkg/common"
)

func main() {
	e := echo.New()
	e.Use(common.HTTPDBClientMongo())

	repository := mongodb.NewUrlRepository()
	logic := logic.NewUrlLogic(repository)

	http.NewHttpHandlers(e, logic)

	e.Start(":2345")
}
