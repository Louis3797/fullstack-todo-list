package todo

import (
	"net/http"

	"github.com/Louis3797/fullstack-todo-list/server/db"
	"github.com/Louis3797/fullstack-todo-list/server/utils"
	"github.com/gin-gonic/gin"
)

func HandleGetAllTodos(c *gin.Context) {

	todos := utils.GetTodos()

	c.IndentedJSON(http.StatusOK, todos)
}

func HandleGetTodo(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No id given"})
		return
	}

	key := db.TodoPrefix + id

	val := utils.GetTodo(key)

	c.IndentedJSON(http.StatusOK, val)
}
