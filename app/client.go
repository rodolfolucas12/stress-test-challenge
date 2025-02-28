package app

import (
	"net/http"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

type ClientInterface interface {
	RequestClient() *http.Response
}

func (c Client) RequestClient(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
