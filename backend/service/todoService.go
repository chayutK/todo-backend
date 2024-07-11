package todo

import (
	"fmt"
	"log"

	repository "github.com/chayutK/todo-backend/repository"
	"github.com/gin-gonic/gin"
)

var db = repository.DB

type Todo struct {
	id     int
	title  string
	status string
}

func GetAllHandler(ctx *gin.Context) {
	_, err := db.Query("SELECT id, title, status FROM todos")

	if err != nil {
		log.Println("Error while query todo list", err)
	}

	// for rows.Next() {
	// 	var id int
	// 	var title, status string
	// 	err := rows.Scan(&id, &title, &status)

	// 	if err != nil {
	// 		log.Println("Error while scanning data", err)
	// 	}

	// 	fmt.Println(id, title, status)
	// }
	fmt.Println("test")

}

func InsertHandler(ctx *gin.Context) {
	// q := `INSERT INTO todos (title,status) value($1,$2) RETURNING id`
	todo := Todo{}
	err := ctx.Bind(&todo)

	if err != nil {
		log.Println("Error while binding request body", err)
	}
	fmt.Println(todo)

}
