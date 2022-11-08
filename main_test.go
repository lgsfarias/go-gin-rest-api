package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/controllers"
	"github.com/lgsfarias/go-gin-rest-api/database"
	"github.com/lgsfarias/go-gin-rest-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CreateStudentMock() *models.Student {
	student := models.Student{
		Name:  "John Doe",
		Email: "john@john.com",
		CPF:   "12345678910",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
	return &student
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCodeHello(t *testing.T) {
	r := RoutesSetup()
	r.GET("/", controllers.Hello)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestVerifyBodyHello(t *testing.T) {
	r := RoutesSetup()
	r.GET("/", controllers.Hello)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	responseMock := `{"message":"Hello World!"}`
	assert.Equal(t, responseMock, resp.Body.String(), "OK response is expected")
}

func TestGetAllStudents(t *testing.T) {
	database.Connect()
	r := RoutesSetup()
	r.GET("/students", controllers.ShowAllStudents)
	req, err := http.NewRequest("GET", "/students", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestGetStudentByCpf(t *testing.T) {
	database.Connect()
	student := CreateStudentMock()
	fmt.Println(student.CPF)
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	req, err := http.NewRequest("GET", "/students/cpf/"+student.CPF, nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Contains(t, resp.Body.String(), student.CPF, "CPF is expected")
}

func TestGetStudentById(t *testing.T) {
	database.Connect()
	student := CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.GET("/students/:id", controllers.GetStudentById)

	req, err := http.NewRequest("GET", "/students/"+fmt.Sprintf("%d", ID), nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	type response struct {
		Data models.Student `json:"data"`
	}
	var respMock response
	err = json.Unmarshal(resp.Body.Bytes(), &respMock)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, ID, int(respMock.Data.ID), "ID is expected")
	assert.Equal(t, student.Name, respMock.Data.Name, "Name is expected")
	assert.Equal(t, student.Email, respMock.Data.Email, "Email is expected")
	assert.Equal(t, student.CPF, respMock.Data.CPF, "CPF is expected")
}

func TestDeleteStudent(t *testing.T) {
	database.Connect()
	CreateStudentMock()

	r := RoutesSetup()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	req, err := http.NewRequest("DELETE", "/students/"+fmt.Sprintf("%d", ID), nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestUpdateStudent(t *testing.T) {
	database.Connect()
	student := CreateStudentMock()
	defer DeleteStudentMock()
	r := RoutesSetup()
	r.PATCH("/students/:id", controllers.UpdateStudent)

	student.Name = "John Doe Updated"
	student.Email = "johnupdated@john.com"
	student.CPF = "12345678911"

	jsonStudent, err := json.Marshal(student)
	if err != nil {
		t.Fatalf("Could not marshal student: %v", err)
	}

	req, err := http.NewRequest("PATCH", "/students/"+fmt.Sprintf("%d", ID), bytes.NewBuffer(jsonStudent))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	type response struct {
		Data models.Student `json:"data"`
	}
	var respMock response

	err = json.Unmarshal(resp.Body.Bytes(), &respMock)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	assert.Equal(t, ID, int(respMock.Data.ID), "ID is expected")
	assert.Equal(t, student.Name, respMock.Data.Name, "Name is expected")
	assert.Equal(t, student.Email, respMock.Data.Email, "Email is expected")
	assert.Equal(t, student.CPF, respMock.Data.CPF, "CPF is expected")
}
