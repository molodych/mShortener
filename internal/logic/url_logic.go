package logic

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"mShorter/internal/app"
	"strings"
)

type urlLogic struct {
	repository app.UrlRepository
}

func NewUrlLogic(repository app.UrlRepository) *urlLogic {
	return &urlLogic{repository: repository}
}

func (u *urlLogic) GetByKey(ctx context.Context, key string) (result map[string]interface{}, err error) {
	resultBSON, err := u.repository.GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	result = make(map[string]interface{})
	result["key"] = resultBSON["key"]
	result["url"] = resultBSON["url"]
	return result, nil
}

func (u *urlLogic) Create(ctx context.Context, url string) (result string, err error) {
	url, err = normalizationUrl(url)
	if err != nil {
		return "", err
	}
	key := generateKey(url)
	urlMap := make(map[string]string)
	urlMap["key"] = key
	urlMap["url"] = url

	err = u.repository.Create(ctx, urlMap)
	if err != nil {
		return "", err
	}
	return key, nil
}

func normalizationUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url is empty")
	}
	if !strings.Contains(url, ".") {
		return "", errors.New("Url is incorrect")
	}
	if !strings.Contains(url, "http://") || !strings.Contains(url, "https://") {
		url = "http://" + url
	}
	return url, nil
}

func generateKey(url string) string {
	h := md5.Sum([]byte(url))
	var key string

	for i, c := range fmt.Sprintf("%x", h) {
		if i < 8 {
			key += string(c)
		} else {
			return key
		}
	}
	return key
}
