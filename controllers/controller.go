package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/database"
	"github.com/lgsfarias/go-gin-rest-api/models"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func ShowAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "No student found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "No student found!"})
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted!"})
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "No student found!"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func GetStudentByCpf(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var student models.Student
	database.DB.Where("cpf = ?", cpf).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "No student found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}
