package main

import (
	"github.com/Louis3797/fullstack-todo-list/server/controller/todo"

	"github.com/Louis3797/fullstack-todo-list/server/db"

	"github.com/gin-gonic/gin"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {

	db.InitRedis()

	r := gin.Default()

	r.GET("/", todo.HandleGetAllTodos)
	r.GET("/todo/:id", todo.HandleGetTodo)
	r.POST("/create-todo", todo.HandleCreateTodo)
	r.PATCH("/todo/:id", todo.HandleUpdateTodo)
	r.DELETE("/delete-todo:id", todo.HandleDeleteTodo)

	r.Run("localhost:4000")

}
