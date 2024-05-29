package main

import (
	"log"
	"net/rpc"

	"github.com/gin-gonic/gin"
	h "h45/handler"
)

func main() {
	router := gin.Default()
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Client connecting error: ", err)
	}

	handler := h.Handler{Client: client}
	router.POST("/user", handler.CreateUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/user/:id", handler.GetUserById)

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Server run error: ", err)
	}
}
