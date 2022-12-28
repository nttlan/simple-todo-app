package main

import (
	"os"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
)

var db *gorm.DB
var err error

// model
type Task struct {
	ID        uint      `json:"id" gorm:"unique, not null"`
	Title     string    `json:"title" gorm:"not null"`
	Status    string    `json:"status" gorm:"type:enum('Not started','Doing','Finished')"`
	Details   string 	`json:"details"`
}

// always called, regardless of being called in main() or not
func init() {
	// load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to mysql database
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	dsn := db_username + ":" + db_password + "@tcp(127.0.0.1:3306)/" + db_database + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	} else {
		log.Println("Connected to MySQL:", db)
	}

	db.AutoMigrate(&Task{})
}

func main() {
	// create a gin router with default middleware
	router := gin.Default()

	router.GET("/tasks", getTasksIndex)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:id", getSingleTask)
	router.PUT("/tasks/:id", updateSingleTask)
	router.DELETE("tasks/:id", deleteSingleTask)

	// by default it serves on :8080 unless a PORT environment variable was defined
	router.Run()
}

func getTasksIndex(c *gin.Context) {
	// get the tasks
	var tasks []Task
	err = db.Find(&tasks).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})	
}


func createTask(c *gin.Context) {
	// get data from request body
	var body struct {
		Title    string `json:"title" binding:"required"`
		Status   string `json:"status" binding:"required"`
		Details  string `json:"details" binding:"required"`
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create a task
	task := Task{
		Title: body.Title,
		Status: body.Status,
		Details: body.Details,
	}
	err = db.Create(&task).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully",
	})
}

func getSingleTask(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	var task Task
	err = db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func updateSingleTask(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from request body
	var body struct {
		Title    string `json:"title" binding:"required"`
		Status   string `json:"status" binding:"required"`
		Details  string `json:"details" binding:"required"`
	}

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// find and update the task
	var task Task
	err = db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.Model(&task).Updates(Task{
		Title: body.Title,
		Status: body.Status,
		Details: body.Details,
		}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}

func deleteSingleTask(c *gin.Context) {
	// get id from url
	id := c.Param("id")


	// delete the tasks
	err = db.Delete(&Task{}, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}