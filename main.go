package main

import (
	"github.com/gin-gonic/gin"
)

func showAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"students": []string{"John", "Paul", "George", "Ringo"},
	})
}

func main() {
	r := gin.Default()
	r.GET("/students", showAllStudents)
	r.Run(":8080")
}
