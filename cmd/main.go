package main

import (
	"github.com/NikCool98/todo-lists"
	"github.com/NikCool98/todo-lists/pkg/handler"
	"github.com/NikCool98/todo-lists/pkg/repository"
	"github.com/NikCool98/todo-lists/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
