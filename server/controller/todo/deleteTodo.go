package todo

import (
	"net/http"

	"github.com/Louis3797/fullstack-todo-list/server/db"
	"github.com/gin-gonic/gin"
)

func HandleDeleteTodo(c *gin.Context) {

	id := c.Param("id")

	key := db.TodoPrefix + id

	err := db.RedisClient.Del(c, key).Err()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not OK"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})

}
