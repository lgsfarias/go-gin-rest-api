package main

import (
	"github.com/lgsfarias/go-gin-rest-api/database"
	"github.com/lgsfarias/go-gin-rest-api/models"
	"github.com/lgsfarias/go-gin-rest-api/routes"
)

func main() {
	database.Connect()
	models.Students = []models.Student{
		{Name: "John", CPF: "123456789", Email: "john.dev@john.com"},
		{Name: "Paul", CPF: "987654321", Email: "paul.dev@paul.com"},
		{Name: "George", CPF: "123456789", Email: "ana.dev@ana.com"},
	}
	routes.HandleRequests()

}
