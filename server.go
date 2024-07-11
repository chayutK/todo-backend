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
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()
	r.GET("/", helloWorldHandler)

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

	db = repository.DatabaseSync()
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
