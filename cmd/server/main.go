package main

import (
	"log"
	"net/http"
	"go18_api/internal/handlers"
	"go18_api/internal/storage"
)

func main() {
	var store storage.Storage = storage.NewStorageData()

	h := handlers.New(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", h.TasksCollection) // GET, POST
	mux.HandleFunc("/tasks/", h.TaskItem)       // GET, PUT, DELETE

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}