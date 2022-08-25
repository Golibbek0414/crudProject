package main

import (
	"context"
	"github.com/exampleEP/crudProject/api/controller"
	"github.com/exampleEP/crudProject/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	root := mux.NewRouter()

	root.HandleFunc("/user", controller.Sign)

	controller.Init(root)

	log.Println(" main: Project is started on the port: ", cfg.HTTPPort)

	errChan := make(chan error, 1)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	httpServer := http.Server{
		Addr:    cfg.HTTPPort,
		Handler: root,
	}

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-errChan:
		log.Println(err)
	case <-osSignals:
		log.Println("main: recieved os signal, shutting down")
		_ = httpServer.Shutdown(context.Background())
		return
	}
}


