package main

// https://golang.cafe/blog/golang-httptest-example.html

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/upper", UpperCaseHandler)
	log.Fatal(http.ListenAndServe(":5555", nil))
}

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func UpperCaseHandler(wr http.ResponseWriter, req *http.Request) {
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		log.Println(err)
		return
	}
	word := q.Get("word")
	if len(word) == 0 {
		wr.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(wr, "missing word")
		return
	}
	wr.WriteHeader(http.StatusOK)
	fmt.Fprint(wr, strings.ToUpper(word))
}
