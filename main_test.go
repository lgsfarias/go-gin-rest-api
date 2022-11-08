package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lgsfarias/go-gin-rest-api/controllers"
	"github.com/stretchr/testify/assert"
)

func RoutesSetup() *gin.Engine {
	r := gin.Default()
	return r
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
