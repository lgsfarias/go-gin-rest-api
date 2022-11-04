package main

import (
	"github.com/lgsfarias/go-gin-rest-api/database"
	"github.com/lgsfarias/go-gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
