package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/database"
	"github.com/lgsfarias/go-gin-rest-api/models"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)

}

func Greet(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}
