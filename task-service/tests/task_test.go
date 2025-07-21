package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"task-service/database"
	"task-service/models"
	"task-service/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router http.Handler

func TestMain(m *testing.M) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "root")
	os.Setenv("DB_HOST", "mysql")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "taskdb")

	fmt.Println("Connexion à la base de données...")
	database.Connect()
	database.DB.AutoMigrate(&models.Task{})

	router = routes.SetupRoutes()
	fmt.Println("Serveur prêt, début des tests...")

	os.Exit(m.Run())
}

func TestFullTaskLifecycle(t *testing.T) {
	fmt.Println("Test 1 - Création de tâche")

	// 1. Création
	task := map[string]interface{}{
		"title":     "Test intégration",
		"completed": false,
	}
	taskJSON, _ := json.Marshal(task)

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var created models.Task
	json.NewDecoder(resp.Body).Decode(&created)
	assert.Equal(t, "Test intégration", created.Title)
	assert.False(t, created.Completed)
	fmt.Printf("Tâche créée avec ID %d\n", created.ID)

	// 2. Lecture
	fmt.Println("Test 2 - Lecture des tâches")
	req = httptest.NewRequest("GET", "/tasks", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "Test intégration")
	fmt.Println("Lecture OK")

	// 3. PUT
	fmt.Println("Test 3 - Modification complète de la tâche")
	updated := map[string]interface{}{
		"title":     "Titre modifié",
		"completed": true,
	}
	updatedJSON, _ := json.Marshal(updated)

	req = httptest.NewRequest("PUT", fmt.Sprintf("/tasks/%d", created.ID), bytes.NewBuffer(updatedJSON))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println("PUT OK")

	// 4. PATCH
	fmt.Println("Test 4 - PATCH (completed=false)")
	patch := map[string]interface{}{
		"completed": false,
	}
	patchJSON, _ := json.Marshal(patch)

	req = httptest.NewRequest("PATCH", fmt.Sprintf("/tasks/%d", created.ID), bytes.NewBuffer(patchJSON))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println("PATCH OK")

	// 5. DELETE
	fmt.Println("Test 5 - Suppression")
	req = httptest.NewRequest("DELETE", fmt.Sprintf("/tasks/%d", created.ID), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNoContent, resp.Code)
	fmt.Println("Suppression OK")

	// 6. Vérification post-suppression
	fmt.Println("Test 6 - Vérification que la tâche a bien disparu")
	req = httptest.NewRequest("GET", "/tasks", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	body, _ = io.ReadAll(resp.Body)
	assert.NotContains(t, string(body), "Titre modifié")
	fmt.Println("La tâche n’est plus présente (test OK)")
}
