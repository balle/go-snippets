package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8888", nil)
}
