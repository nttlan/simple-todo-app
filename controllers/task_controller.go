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

func (c* TaskController) GetTasks(context *gin.Context) {
	tasks, err := c.service.FindAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (c* TaskController) CreateTask(context *gin.Context) {
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

	result, err := c.service.Create(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, result)
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
	context.JSON(http.StatusOK, task)
}