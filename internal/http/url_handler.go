package http

import (
	"github.com/labstack/echo/v4"
	"mShorter/internal/app"
	"net/http"
)

type httpHandlers struct {
	urlLogic app.UrlLogic
}

func NewHttpHandlers(e *echo.Echo, urlLogic app.UrlLogic) {
	h := httpHandlers{urlLogic: urlLogic}
	e.GET("/s/:key", h.GetByKey)
	e.POST("/a/", h.Create)
}

func (h *httpHandlers) GetByKey(c echo.Context) error {
	ctx := c.Request().Context()
	key := c.Param("key")
	result, err := h.urlLogic.GetByKey(ctx, key)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	return c.Redirect(http.StatusFound, result["url"].(string))
}

func (h *httpHandlers) Create(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.urlLogic.Create(ctx, c.QueryParam("url"))
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	return c.JSON(http.StatusCreated, result)
}
