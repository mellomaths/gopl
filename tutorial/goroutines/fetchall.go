// Fetchall searchs for URLs concurrently and informs the payload size and how much time took to response
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, channel chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err) // sends error to channel
		return
	}

	size, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close() // avoids leak of resources
	if err != nil {
		channel <- fmt.Sprintf("while reading %s: %v", url, err) // sends reading error to the channel
		return
	}

	secs := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", secs, size, url) // sends to channel the time and size of the request
}

func main() {
	start := time.Now()
	channel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, channel) // starts a go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-channel) // reads the channel information
	}

	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}
