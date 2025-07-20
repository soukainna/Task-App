package main

import (
	"log"
	"net/http"
	"os"
	"task-service/database"
	"task-service/models"
	"task-service/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Task{})

	port := getEnv("PORT", "8080")
	log.Println("🚀 Serveur démarré sur le port", port)
	log.Fatal(http.ListenAndServe(":"+port, routes.SetupRoutes()))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
