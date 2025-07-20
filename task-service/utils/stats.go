package utils

import (
	"log"
	"sync"
	"sync/atomic"
	"task-service/database"
	"task-service/models"
)

var completedCount int64

func CountCompletedTasks(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int64
	var tasks []models.Task
	database.DB.Find(&tasks)
	for _, t := range tasks {
		if t.Completed {
			count++
		}
	}
	atomic.StoreInt64(&completedCount, count)
	log.Println("📊 Nombre de tâches complétées :", count)
}

func GetCompletedCount() int64 {
	return atomic.LoadInt64(&completedCount)
}
