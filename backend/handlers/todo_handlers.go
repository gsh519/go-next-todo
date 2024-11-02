package handlers

import (
	"database/sql"
	"go-todo/db"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func bindAndValidate(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		return err
	}
	return validate.Struct(data)
}

func GetTodos(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := db.FetchTodos(database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodoHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var contentData models.ContentData
		if err := bindAndValidate(c, &contentData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		id, err := db.CreateTodo(database, contentData.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func UpdateTodoHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("todoId")
		var contentData models.ContentData
		if err := bindAndValidate(c, &contentData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := db.UpdateTodoContent(database, todoId, contentData.Content); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func DeleteTodoHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("todoId")
		if err := db.DeleteTodoByID(database, todoId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func DoneTodoHandler(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId := c.Param("todoId")
		if err := db.MarkTodoAsDone(database, todoId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
