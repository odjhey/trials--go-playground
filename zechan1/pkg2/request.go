package pkg2

import (
	"fmt"
	"io"
	"net/http"
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

// A Fan-Out?
// Why? Fan out the tasks to multiple goroutines?
func Run() {
	fmt.Println("hello pkg2", urls)
	c := make(chan response, len(urls))

	// run all requests
	for _, v := range urls {
		fmt.Println("queing: ", v)
		// send
		go func(url string) {
			fmt.Println("running url: ", url)
			r, err := http.Get(url)
			c <- response{url: url, response: r, err: err}

		}(v)

	}

	// receive
	for i := 0; i < 3; i++ {
		r := <-c
		if r.err != nil {
			println("response: error", r.url, r.err.Error())
			continue
		}

		b, err := io.ReadAll(r.response.Body)
		_ = b
		if err != nil {
			println("response: error", r.url, err)
			continue
		}
		defer r.response.Body.Close()

		println("response: ", r.response.Status, r.url)
	}

}
