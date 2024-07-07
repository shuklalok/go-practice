package main

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c *Client) GetUpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", fmt.Errorf("unable to complete GET request: %v", err)
	}
	defer res.Body.Close()
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response: %v", err)
	}
	return string(out), nil
}
