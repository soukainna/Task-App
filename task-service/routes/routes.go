package routes

import (
	"net/http"
	"task-service/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controllers.PatchTask).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/stats", controllers.GetStats).Methods("GET")

	return r
}
