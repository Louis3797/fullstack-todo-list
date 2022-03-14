package todo

import (
	"encoding/json"
	"net/http"

	"github.com/Louis3797/fullstack-todo-list/server/db"
	m "github.com/Louis3797/fullstack-todo-list/server/models"
	"github.com/gin-gonic/gin"
)

func HandleUpdateTodo(c *gin.Context) {

	var newTodo m.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}

	key := db.TodoPrefix + newTodo.ID

	err = db.RedisClient.Set(c, key, bytes, 0).Err()

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newTodo)

}
