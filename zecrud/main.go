package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zecrud/controllers"
	"zecrud/database"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")

	fmt.Println("ei start!")

	err := start()
	if err != nil {
		panic(err)
	}

}

func start() error {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	mux := http.NewServeMux()
	mux.HandleFunc("/food/", controllers.Food)

	ctx, _ := context.WithCancel(context.Background())
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 8 * time.Second,
		Handler:      mux,
		BaseContext: func(l net.Listener) context.Context {
			db, _ := database.Open()
			ctx = context.WithValue(ctx, "dbKey", db)
			return ctx
		},
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
