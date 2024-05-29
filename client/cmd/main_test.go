package main

import (
	"bytes"
	"encoding/json"
	h "h45/handler"
	"net/http"
	"net/http/httptest"
	"net/rpc"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setRouter(client *rpc.Client) *gin.Engine {
	router := gin.Default()

	handler := h.Handler{Client: client}

	router.POST("/user", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/user/:id", handler.GetUserById)

	return router
}

func performRequest(router http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestCreateUser(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		t.Fatal("Client connecting error: ", err)
	}
	defer client.Close()

	router := setRouter(client)

	user := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "john@example.com",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	w := performRequest(router, "POST", "/user", userJSON)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		t.Fatal("Client connecting error: ", err)
	}
	defer client.Close()

	router := setRouter(client)

	w := performRequest(router, "GET", "/users", nil)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserById(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		t.Fatal("Client connecting error: ", err)
	}
	defer client.Close()

	router := setRouter(client)

	w := performRequest(router, "GET", "/user/1", nil)

	assert.Equal(t, http.StatusOK, w.Code)
}
