package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	TodoId    int
	Content   string
	IsDone    bool
	DeletedAt sql.NullTime
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	// 接続の確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("データベース接続成功")

	return db
}

func fetchTodos(db *sql.DB) *sql.Rows {
	rows, err := db.Query("select * from todos where deleted_at is null")
	if err != nil {
		fmt.Println("エラー")
		panic(err.Error())
	}

	return rows
}

func findTodo(db *sql.DB, todoId string) *sql.Row {
	row := db.QueryRow("select * from todos where deleted_at is null and todo_id = ? limit 1", todoId)

	return row
}

func createTodo(db *sql.DB, content string) {
	result, err := db.Exec("insert into todos (content) values (?)", content)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(id)
}

func updateTodo(db *sql.DB, todoId string, content string) {
	_, err := db.Exec("update todos set content = ? where todo_id = ?", content, todoId)

	if err != nil {
		log.Fatalln(err)
	}
}

func deleteTodo(db *sql.DB, todoId string) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	now := time.Now().In(jst)

	_, err = db.Exec("update todos set deleted_at = ? where todo_id = ?", now, todoId)

	if err != nil {
		log.Fatalln(err)
	}
}

func doneTodo(db *sql.DB, todoId string) {
	_, err := db.Exec("update todos set is_done = 1 where todo_id = ?", todoId)

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	db := connectDB()
	defer db.Close()

	router := gin.Default()

	// 一覧取得
	router.GET("/todos", func(c *gin.Context) {
		todos := []Todo{}
		todo := Todo{}

		rows := fetchTodos(db)

		for rows.Next() {
			error := rows.Scan(&todo.TodoId, &todo.Content, &todo.IsDone, &todo.DeletedAt)
			if error != nil {

				fmt.Println("scan error")
				log.Fatal(error)
			} else {
				todos = append(todos, todo)
			}
		}

		c.JSON(http.StatusOK, todos)
	})

	// 登録
	router.POST("/todo", func(c *gin.Context) {
		var json struct {
			Content string `json:"content"`
		}

		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		}

		createTodo(db, json.Content)
	})

	// 更新
	router.PUT("/todo/:todoId", func(c *gin.Context) {
		todoId := c.Param("todoId")

		var param struct {
			Content string `json:"Content"`
		}

		if err := c.BindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invallid Parameter"})
		}

		updateTodo(db, todoId, param.Content)
	})

	// 削除
	router.DELETE("/todo/:todoId", func(c *gin.Context) {
		todoId := c.Param("todoId")

		deleteTodo(db, todoId)
	})

	// 完了
	router.PUT("/todo/:todoId/done", func(c *gin.Context) {
		todoId := c.Param("todoId")

		doneTodo(db, todoId)
	})

	router.Run(":8080")
}
