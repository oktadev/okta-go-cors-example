package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var messages []string

func GetMessages(c *gin.Context) {
	version := c.Param("version")
	if version == "v2" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func PutMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	text := c.PostForm("puttext")
	messages[id] = text
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Credentials", "true")
	version := c.Param("version")
	if version == "v2" {
		c.Header("Access-Control-Expose-Headers", "X-Custom")
	}
	c.Header("X-Custom", "123456789")
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func OptionMessage(c *gin.Context) {
	for key, value := range c.Request.Header {
		fmt.Println(key, ": ", value)
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
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
