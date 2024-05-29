package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var (
	users []User
)

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	var userServer User

	rpc.Register(&userServer)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}
	fmt.Println("Server is listenig on :8080")
	http.Serve(listener, nil)
}

func (u *User) CreateUser(arg *User, res *User) error {
	if arg == nil {
		return errors.New("Nill data from argument")
	}
	users = append(users, *arg)
	*res = *arg
	return nil
}

func (u *User) GetUsers(arg *string, res *[]User) error {
	*res = users
	return nil
}

func (u *User) GetUserById(id *int, res *User) error {
	for _, v := range users {
		if v.Id == *id {
			*res = v
			return nil
		}
	}
	return errors.New("User not found!")
}
