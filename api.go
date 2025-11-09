package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// POST  /tasks
func createTaskHandler(c *gin.Context) {
	var input struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.shouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return 
	}

	id, err := createTask(input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return 
	}

	task := Task{
		ID: id, 
		Description: input.Description,
		Status: "pending",
		CreatedAt: time.Now(), 
	}

	// Enqueue task asynchronously
	go enqueTask(taskQueue, task)
	
	c.JSON(http.StatusCreated, task)

}


// GET /tasks
func getAllTasksHandler(c *gin.Context) {
	tasks, err := getAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}


// GET /tasks/:id
func getTaskByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := getTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)  
}