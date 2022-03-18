package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
	"testing"
)

type logic struct{}

func (*logic) GetByKey(ctx context.Context, key string) (result map[string]interface{}, err error) {
	if key == "12345678" {
		return map[string]interface{}{"key": "12345678", "url": "http://w3.org"}, nil
	}
	return nil, errors.New("We have a problem")
}
func (*logic) Create(ctx context.Context, url string) (result string, err error) {
	if url == "http://w3.org" {
		return "12345678", nil
	}
	return "", errors.New("We have a problem")
}

func TestHandler(t *testing.T) {
	go func() {
		e := echo.New()
		NewHttpHandlers(e, &logic{})
		e.Start(":2345")
	}()
	client := &http.Client{}
	t.Run("TestGetByKey positive", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://localhost:2345/s/12345678", nil)
		if err != nil {
			t.Error(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		defer res.Body.Close()

		if res.Status != "200 OK" {
			t.Errorf("Different code. Expected: 200 OK, in fact: %s", res.Status)
		}
	})
	t.Run("TestGetByKey negative", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://localhost:2345/s/87654321", nil)
		if err != nil {
			t.Error(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		defer res.Body.Close()

		if res.Status != "502 Bad Gateway" {
			t.Errorf("Different code. Expected: 502 Bad Gateway, in fact: %s", res.Status)
		}
	})
	t.Run("Create positive", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:2345/a/?url=http://w3.org", nil)
		if err != nil {
			t.Error(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		defer res.Body.Close()

		if res.Status != "201 Created" {
			t.Errorf("Different code. Exepcted code: 201 Created, in fact: %s", res.Status)
		}
	})
	t.Run("Create negative", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:2345/a/?url=http://w4.org", nil)
		if err != nil {
			t.Error(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			t.Error(err)
			return
		}
		defer res.Body.Close()

		if res.Status != "502 Bad Gateway" {
			t.Errorf("Different code. Exepcted code: 502 Bad Gateway, in fact: %s", res.Status)
		}
	})
}
