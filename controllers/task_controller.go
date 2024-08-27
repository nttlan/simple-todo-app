package controllers

import (
	"net/http"

	"github.com/emerald-lan/simple-todo-app/models"
	"github.com/emerald-lan/simple-todo-app/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{service}
}

func (c *TaskController) GetTasks(context *gin.Context) {
	tasks, err := c.service.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Tasks fetched successfully",
		"tasks":   tasks,
	})
}

func (c *TaskController) CreateTask(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if task.Title == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Task title is required"})
		return
	}

	task.Completed = false // default

	_, err = c.service.Create(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
	})
}

func (c *TaskController) GetTaskById(context *gin.Context) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := c.service.FindById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Task fetched successfully",
		"task":    task,
	})
}

func (c *TaskController) UpdateTaskById(context *gin.Context) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	err = context.ShouldBindJSON(&updatedTask)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.service.FindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	task.Completed = updatedTask.Completed

	_, err = c.service.Update(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":     "Task updated successfully",
		"updatedTask": task,
	})
}

func (c *TaskController) DeleteTaskById(context *gin.Context) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err = c.service.FindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	_, err = c.service.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}