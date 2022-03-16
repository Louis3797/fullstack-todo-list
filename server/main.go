package main

import (
	"github.com/Louis3797/fullstack-todo-list/server/controller/todo"
	"github.com/gin-contrib/cors"

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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))

	r.GET("/", todo.HandleGetAllTodos)
	r.GET("/todo/:id", todo.HandleGetTodo)
	r.POST("/create-todo", todo.HandleCreateTodo)
	r.PATCH("/todo/:id", todo.HandleUpdateTodo)
	r.DELETE("/delete-todo/:id", todo.HandleDeleteTodo)

	r.Run("localhost:4000")

}
