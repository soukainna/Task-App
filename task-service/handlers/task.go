package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"task-service/models"
)

func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, done FROM tasks")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []models.Task
		for rows.Next() {
			var t models.Task
			if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, t)
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t models.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow("INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id", t.Title, t.Done).Scan(&t.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(t)
	}
}
