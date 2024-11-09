package main

import (
	"go-todo/db"
	"go-todo/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// タイムゾーン設定
	_ = os.Setenv("TZ", "Asia/Tokyo")
	log.Println("init")

	// データベース接続
	database := db.ConnectDB()
	defer database.Close()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/todos", handlers.GetTodos(database))
	router.POST("/todo", handlers.CreateTodoHandler(database))
	router.PUT("/todo/:todoId", handlers.UpdateTodoHandler(database))
	router.DELETE("/todo/:todoId", handlers.DeleteTodoHandler(database))
	router.PUT("/todo/:todoId/done", handlers.DoneTodoHandler(database))

	router.Run(":8080")
}
