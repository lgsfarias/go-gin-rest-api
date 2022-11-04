package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/models"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greet(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}
