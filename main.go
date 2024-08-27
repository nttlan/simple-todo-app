package main

import (
    "log"
	"github.com/gin-gonic/gin"
    "github.com/emerald-lan/simple-todo-app/config"
    "github.com/emerald-lan/simple-todo-app/controllers"
    "github.com/emerald-lan/simple-todo-app/services"
)

func main() {
    client := config.InitDB()
    defer config.DisconnectDB(client)

    db := client.Database("simple-todo-app")

    taskService := services.NewTaskService(db)
    
    taskController := controllers.NewTaskController(taskService)

	router := gin.Default()
    log.Println("Server started on port 8080")

    router.POST("/tasks", taskController.CreateTask)
    router.GET("/tasks", taskController.GetTasks)
    router.GET("/tasks/:id", taskController.GetTaskById)

	router.Run(":8080")
}