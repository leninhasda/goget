package goget

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(c Client) *getter {
	return &getter{client: c}
}

type getter struct {
	client Client
}

func (g *getter) GetMd5Hash(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hash := md5.Sum(body)
	return hex.EncodeToString(hash[:]), nil
}
