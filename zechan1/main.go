package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type response struct {
	url      string
	response *http.Response
	err      error
}

var urls = []string{
	"https://github.com",
	"https://google.com",
	"https://somexaklsj.com",
}

var wg = &sync.WaitGroup{}

func main() {
	fmt.Println("hello world", urls)
	c := make(chan response, len(urls))

	wg.Add(1)
	// run all requests
	go func() {
		defer wg.Done()

		for _, v := range urls {
			wg.Add(1)
			// send
			go func(url string) {
				defer wg.Done()

				fmt.Println("running url: ", url)
				r, err := http.Get(url)
				c <- response{url: url, response: r, err: err}

			}(v)

			wg.Add(1)
			// receive
			go func() {
				defer wg.Done()
				r := <-c
				if r.err != nil {
					println("response: error", r.url, r.err.Error())
					return
				}

				b, err := io.ReadAll(r.response.Body)
				_ = b
				if err != nil {
					println("response: error", r.url, err)
					return
				}
				defer r.response.Body.Close()

				println("response: ", r.response.Status, r.url)
			}()

		}

	}()

	wg.Wait()

}
