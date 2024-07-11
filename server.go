package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repository "github.com/chayutK/todo-backend/repository"
	// todo "github.com/chayutK/todo-backend/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repository.Sync()
	db = repository.DB
	defer db.Close()

	r := gin.Default()

	r.GET("/", helloWorldHandler)
	r.GET("/api/v1/todos", GetAllHandler)
	r.GET("/api/v1/todos/:id", GetHandler)

	r.POST("/api/v1/todos", InsertHandler)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closeChannel := make(chan struct{})
	go func() {
		<-ctx.Done()
		fmt.Println("Server is shutting down.....")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err.Error())
			}
		}
		close(closeChannel)

	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}

	<-closeChannel
	fmt.Println("Server is closed")

}

func helloWorldHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Hello, World",
	})
}

// func DatabaseSync() *sql.DB {
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URI"))

// 	if err != nil {
// 		log.Fatal("Connection to database error", err)
// 	}

// 	fmt.Println("Database sync successfully")
// 	fmt.Println(db)
// 	return db
// }

type Todo struct {
	ID     int `url:"id"`
	Title  string
	Status string
}

func GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	q := "SELECT id,title,status FROM todos WHERE id=$1"
	row := db.QueryRow(q, id)

	var todo Todo

	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)

	if err != nil {
		log.Println("Error while scanning data", err)
	}

	ctx.JSON(http.StatusOK, todo)

}

func GetAllHandler(ctx *gin.Context) {
	rows, err := repository.DB.Query("SELECT id, title, status FROM todos")

	if err != nil {
		log.Println("Error while query todo list", err)
	}

	result := []Todo{}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status)

		if err != nil {
			log.Println("Error while scanning data", err)
		}

		result = append(result, todo)
	}

	ctx.JSON(200, result)
}

func InsertHandler(ctx *gin.Context) {

	var todo Todo
	err := ctx.ShouldBindJSON(&todo)

	if err != nil {
		log.Println("Error while binding request body", err)
	}
	q := `INSERT INTO todos (title,status) values ($1,$2) RETURNING id,title,status`
	row := db.QueryRow(q, todo.Title, todo.Status)

	var result Todo

	err = row.Scan(&result.ID, &result.Title, &result.Status)
	if err != nil {
		log.Println("Error while scan id", err)
	}
	ctx.JSON(http.StatusOK, result)

}
