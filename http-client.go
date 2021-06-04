package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8888/hello?name=balle")

	if err != nil {
		log.Panicf("Cannot get url: %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Panicf("Cannot read body: %v", err)
	}

	fmt.Printf("Got %d\n%s\n", res.StatusCode, body)
}
