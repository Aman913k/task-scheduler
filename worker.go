package main

import "time"

func startWorker(queue TaskQueue) {
	for task := range queue {
		log.Printf("Processing task %d: %s", task.ID, task.Description)

		if err := updateTaskStatus(task.ID, "running"); err != nil {
			log.Printf("Failed to update task %d status: %v", task.ID, err)
			continue
		}

		// Simulate task processing
		time.Sleep(2 * time.Second)

		if err := updateTaskStatus(task.ID, "completed"); err != nil {
			log.Printf("Failed to update task %d status: %v", task.ID, err)
		}
		log.Printf("Completed task %d", task.ID)
	}
}