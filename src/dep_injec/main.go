package main

import (
	"fmt"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

type UserService struct {
	db *Database
}

func NewUserService(db *Database) *UserService {
	return &UserService{db: db}
}

func main() {
	db := NewDatabase()
	userService := NewUserService(db)

	fmt.Println("UserService initialized with Database:", userService.db)
}
