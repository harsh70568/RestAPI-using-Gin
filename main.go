package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"Item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: true},
	{ID: "2", Item: "Drink Water", Completed: false},
	{ID: "3", Item: "Read Book", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoID(context *gin.Context) {
	id := context.Param("id")
	for _, t := range todos {
		if t.ID == id {
			context.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
}

func main() {
	router := gin.Default()

	router.GET("/getTodos", getTodos)
	router.GET("/Todos/:id", getTodoID)
	router.POST("/sendTodos", addTodos)

	router.Run("localhost:9090")
}
