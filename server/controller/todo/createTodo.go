package todo

import (
	"encoding/json"
	"net/http"

	"github.com/Louis3797/fullstack-todo-list/server/db"
	m "github.com/Louis3797/fullstack-todo-list/server/models"
	"github.com/gin-gonic/gin"
)

func HandleCreateTodo(c *gin.Context) {

	var newTodo m.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	bytes, err := json.Marshal(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}

	redisErr := db.RedisClient.Set(c, (db.TodoPrefix + newTodo.ID), bytes, 0).Err()

	if redisErr != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newTodo)
}
