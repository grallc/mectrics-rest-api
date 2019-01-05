package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Connection - Model (struct)
type Connection struct {
	ID          string  `json:"id"`
	ServerURL   string  `json:"server"`
	InitialTime float64 `json:"initial"`
	FinalTime   float64 `json:"final"`
	Duration    float64 `json:"duration"`
}

// GetConnectionsEndpoint - GET all connections
func GetConnectionsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

// AddConnectionsEndpoint - POST a new connection
func AddConnectionsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	fmt.Println("Starting server on port 8001...")
	r := mux.NewRouter()
	r.HandleFunc("/metrics", GetConnectionsEndpoint).Methods("GET")
	r.HandleFunc("/metrics", AddConnectionsEndpoint).Methods("POST")
	if err := http.ListenAndServe(":8001", r); err != nil {
		log.Fatal(err)
	}
}
