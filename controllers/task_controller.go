package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/emerald-lan/simple-todo-app/models"
	"github.com/emerald-lan/simple-todo-app/services"
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

	result, err := c.service.CreateOne(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, result)
}