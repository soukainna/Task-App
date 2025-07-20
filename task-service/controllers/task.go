package controllers

import (
	"encoding/json"
	"net/http"
	"sync"
	"task-service/database"
	"task-service/models"
	"task-service/utils"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go utils.CountCompletedTasks(&wg)

	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)

	wg.Wait()
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	count := utils.GetCompletedCount()
	json.NewEncoder(w).Encode(map[string]int64{"completed_tasks": count})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func PatchTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	var patch struct {
		Completed *bool `json:"completed"`
	}
	json.NewDecoder(r.Body).Decode(&patch)
	if patch.Completed != nil {
		task.Completed = *patch.Completed
		database.DB.Save(&task)
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Delete(&models.Task{}, id)
	w.WriteHeader(http.StatusNoContent)
}
