package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Connection - Model (struct)
type Connection struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	ServerURL   string        `json:"server"`
	InitialTime float64       `json:"initial"`
	FinalTime   float64       `json:"final"`
	Duration    float64       `json:"duration"`
}

var (
	session    *mgo.Session
	collection *mgo.Collection
)

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
	log.Println("Starting server on port 8001...")
	r := mux.NewRouter()
	r.HandleFunc("/metrics", GetConnectionsEndpoint).Methods("GET")
	r.HandleFunc("/metrics", AddConnectionsEndpoint).Methods("POST")

	log.Println("Starting MongoDB session")
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("connectionsdb").C("connections")

	if err := http.ListenAndServe(":8001", r); err != nil {
		log.Fatal(err)
	}
}
