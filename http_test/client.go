package main

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
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
		return "", errors.Wrap(err, "unable to complete GET request")
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "unable to read response")
	}
	return string(out), nil
}
