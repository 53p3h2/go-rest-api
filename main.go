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

func main() {
	router := gin.Default()
	router.GET("/", Hi)
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTasks)
	router.PUT("/tasks/:id", updateTasks)
	router.DELETE("/tasks/:id", deleteTasks)
	router.Run("localhost:3001")
}
