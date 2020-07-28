package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"aviasales/src/cities"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	cities.CityRegister(router.Group(""))
	cities.InitPeriodicalTask()
	cities.RunTasks()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, os.Interrupt)
	<-quit

	fmt.Println("Shutdown Server ...")
}
