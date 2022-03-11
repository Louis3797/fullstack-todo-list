package utils

import (
	"encoding/json"

	"github.com/Louis3797/fullstack-todo-list/server/db"
	m "github.com/Louis3797/fullstack-todo-list/server/models"
	"github.com/go-redis/redis/v8"
)

func GetTodos() (todos []m.Todo) {
	keys, err := db.RedisClient.Keys(db.Ctx, "*").Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		todo := GetTodo(key)
		todos = append(todos, todo)
	}
	return todos
}

func GetTodo(key string) (todo m.Todo) {
	val, err := db.RedisClient.Get(db.Ctx, key).Result()

	if err != nil {
		panic(err)
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(val), &todo)
	}

	return todo
}
