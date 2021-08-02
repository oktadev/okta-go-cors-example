package main

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

var messages []string

func GetMessages(c *gin.Context) {
	version := c.Param("version")
	fmt.Println("Version", version)
	if version == "v2" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func PutMessage(c *gin.Context) {
	version := c.Param("version")
	id, _ := strconv.Atoi(c.Param("id"))
	text := c.PostForm("puttext")
	messages[id] = text
	if version == "v2" {
        c.Header("Access-Control-Expose-Headers", "X-Custom")
	    c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
        c.Header("Access-Control-Allow-Credentials", "true")
	}
	c.Header("X-Custom", "123456789")
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func OptionMessage(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	c.Header("Access-Control-Allow-Headers", "X-Token")
    c.Header("Access-Control-Allow-Credentials", "true")
}

func main() {
	messages = append(messages, "Hello CORS!")
	r := gin.Default()
	r.GET("/api/:version/messages", GetMessages)
	r.PUT("/api/:version/messages/:id", PutMessage)
	r.OPTIONS("/api/v2/messages/:id", OptionMessage)
	r.Run(":8000")
}
