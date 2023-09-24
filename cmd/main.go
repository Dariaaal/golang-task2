package main

import (
	"github.com/dariaaal/golang-task2/internal"
	// "github.com/dariaaal/golang-task2/pkg/handler"
	// "github.com/dariaaal/golang-task2/pkg/repository"
	// "github.com/dariaaal/golang-task2/pkg/service"
)

func main() {
	// repos := repository.NewRepository()
	// services := service.NewService(repos)
	// handlers := handler.NewHandler(services)

	internal.GetRouter()
}
