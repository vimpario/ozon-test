package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"ozon-test/internal/entities"
	"ozon-test/internal/handlers"
	"ozon-test/pkg/helpers"
)

func main() {
	logger := helpers.NewLogger()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	logger.Info("starting")
	if server, err := handlers.NewServer(logger); err != nil {
		logger.Error(err)
	} else if err := server.Run(); err != nil {
		logger.Fatal(entities.RunError)
	}
}
