package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8001...")
	http.ListenAndServe(":8001", nil)

}
