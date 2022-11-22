package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Sutartu")

	err := start()
	if err != nil {
		fmt.Println("Eyyy ", err)
	}
}

func start() error {

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	waiters := make(chan int, 1)

	go func() {
		time.Sleep(5 * time.Second)
		waiters <- 1
	}()

	select {
	case <-waiters:
		fmt.Println("Sleep donoers")
		return nil
	case <-shutdown:
		return errors.New("Manual kill!")
	}

}
