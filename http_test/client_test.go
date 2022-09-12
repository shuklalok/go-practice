package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	// A Server is an HTTP server listening on a system-chosen port
	// on the local loopback interface, for use in end-to-end HTTP tests.
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	// Create client by passing the server URL.
	c := NewClient(svr.URL)
	res, err := c.GetUpperCase("anything")
	if err != nil {
		t.Errorf("expected error to be nill got %v", err)
	}

	// res: expected\r\n
	// due to the http protocol cleanup response
	res = strings.TrimSpace(res)
	if res != expected {
		t.Errorf("expected response %s got %s", expected, res)
	}

}
