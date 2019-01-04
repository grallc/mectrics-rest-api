package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	fmt.Println("Starting server on port 8001...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
