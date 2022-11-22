package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zeordermore/controllers"
)

func main() {
	fmt.Println("ei start!")

	err := start()
	if err != nil {
		panic(err)
	}

}

// GET order
// POST new order
// POST new inventory
// GET inventory
func start() error {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	mux := http.NewServeMux()
	mux.HandleFunc("/order", controllers.CreateOrder)

	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 8 * time.Second,
		Handler:      mux,
	}

	serverError := make(chan error, 1)
	go func() {

		serverError <- api.ListenAndServe()

	}()

	select {
	case sig := <-shutdown:
		fmt.Println("Shutting down with signal", sig)
		// TODO: add a "drain" activity before killing
		return nil
	case err := <-serverError:
		return err
	}

}
