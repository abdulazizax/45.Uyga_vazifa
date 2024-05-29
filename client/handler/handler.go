package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type Handler struct {
	Client *rpc.Client
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var body User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	res, err := h.CreateUserFunk(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	res, err := h.GetUsersFunk("ok")
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get users"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, res)
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	res, err := h.GetUserByIdFunk(newId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, res)
}

func (h *Handler) CreateUserFunk(newUser User) (*User, error) {
	err := h.Client.Call("User.CreateUser", newUser, &newUser)
	if err != nil {
		log.Println("Client invocation error: ", err)
		return nil, err
	}

	fmt.Println("User Registered: ", newUser)
	return &newUser, nil
}

func (h *Handler) GetUsersFunk(str string) (*[]User, error) {
	var users []User
	err := h.Client.Call("User.GetUsers", str, &users)
	if err != nil {
		log.Println("Client invocation error: ", err)
		return nil, err
	}

	fmt.Println("All Users: ", users)
	return &users, nil
}

func (h *Handler) GetUserByIdFunk(id int) (*User, error) {
	var user User
	err := h.Client.Call("User.GetUserById", id, &user)
	if err != nil {
		log.Println("Client invocation error: ", err)
		return nil, err
	}

	fmt.Printf("User : %v whit %v id \n", user, id)
	return &user, nil
}
