package controllers

import "github.com/gin-gonic/gin"

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"students": []string{"John", "Paul", "George", "Ringo"},
	})
}
