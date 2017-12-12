package main

import (
	"github.com/vladborsh/go-rest/handlers"
	"github.com/vladborsh/go-rest/storage"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	db := storage.NewDataBase()
	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))
	log.Printf("serving on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
