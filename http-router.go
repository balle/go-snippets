package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, req *http.Request) {
	user := mux.Vars(req)["user"]
	fmt.Fprintf(w, "hello %s", user)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{user:[a-z]+}", handler)
	http.ListenAndServe(":8888", r)
}
