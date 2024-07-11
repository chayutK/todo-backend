package main

import (
	"database/sql"
	"net/http"

	repository "github.com/chayutK/todo-backend/repository"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {

	db = repository.DatabaseSync()

	r := gin.Default()

	r.GET("/", helloWorldHandler)

	r.Run()

}

func helloWorldHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Hello, World",
	})
}
