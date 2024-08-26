package main

import (
    "log"
	"github.com/gin-gonic/gin"
    "github.com/emerald-lan/simple-todo-app/config"
)

func main() {
    db := config.InitDB()
    defer config.DisconnectDB(db)

	router := gin.Default()
    log.Println("Hello, World!")

	router.Run(":8080")
}