package main

import (
	"log"
	"net/http"
	"task-service/database"
	"task-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db := database.Connect()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.GetTasks(db)).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask(db)).Methods("POST")

	log.Println("Task service listening on port 8080")
	http.ListenAndServe(":8080", r)
}
