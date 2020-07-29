package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"aviasales/src/cities"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var router *gin.Engine

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0777)
		if err != nil {
		}
	}

	file, err := os.OpenFile("logs/travelpayouts.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
		log.WithFields(log.Fields{
			"package": "main",
		}).Info("Initializing logger are successful!")
	} else {
		log.SetOutput(os.Stdout)
		log.Warning("Failed to log to file, using default stderr")
	}
}

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
			log.WithFields(log.Fields{
				"package": "main",
			}).Panic("Web Server can`t start listen!")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, os.Interrupt)
	<-quit
	cities.StopTasks()

	log.WithFields(log.Fields{
		"package": "main",
	}).Info("Initializing logger are successful!")
}
