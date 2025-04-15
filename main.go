package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hi(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hi")
}

func main() {
	router := gin.Default()
	router.GET("/", Hi)
	router.Run("localhost:3001")
}
