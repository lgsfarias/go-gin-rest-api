package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.Run(":8080")
}
