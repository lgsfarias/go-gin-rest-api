package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.Hello)
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.Run(":8080")
}
