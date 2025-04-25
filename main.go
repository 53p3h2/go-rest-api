package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func Hi(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hi")
}

type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Text   string `json:"text"`
	Status bool   `json:"status"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Task{})
}

func main() {
	initDB()
	router := gin.Default()
	router.GET("/", Hi)
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTasks)
	router.PUT("/tasks/:id", updateTasks)
	router.DELETE("/tasks/:id", deleteTasks)
	router.Run("localhost:3001")
}

func getTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func createTasks(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func updateTasks(c *gin.Context) {
	var task Task
	id := c.Param("id")

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Text = input.Text
	task.Status = input.Status
	db.Save(&task)

	c.JSON(http.StatusOK, task)
}

func deleteTasks(c *gin.Context) {
	var task Task
	id := c.Param("id")

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	db.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
