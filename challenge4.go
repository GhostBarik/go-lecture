package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	// list of URLs to fetch
	urls := []string{
		"https://golang.org/",
		"https://golang.org/cmd/",
		"https://golang.org/pkg",
		"https://golang.org/pkg/fmt/",
		"https://golang.org/pkg/os/",
	}

	// SEQUENTIAL VERSION //////////////////////////////

	startTime := time.Now()

	sendRequestsSequentially(urls)

	endTime := time.Now()
	totalTime := endTime.Sub(startTime).Seconds()
	fmt.Printf("sequential version took %v seconds to complete \n", totalTime)
	fmt.Println("*****************************************************")

	// PARALLEL VERSION ////////////////////////////////

	startTime = time.Now()

	// TODO: implement parallel version
	sendRequestsInParallel(urls)

	endTime = time.Now()
	totalTime = endTime.Sub(startTime).Seconds()
	fmt.Printf("parallel version took %v seconds to complete", totalTime)
}

func sendRequestsSequentially(urls []string) {
	for _, url := range urls {
		numberOfBytes := getRequest(url)
		fmt.Printf("get response from %v - %d bytes of data \n", url, numberOfBytes)
	}
}

func sendRequestsInParallel(urls []string) {
	// TODO: implement parallel version
}

func performHttpGetRequest(url string, responseChannel chan int) {
	responseChannel <- getRequest(url)
}

func getRequest(url string) int {

	// create http client
	client := &http.Client{}

	// send GET request
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)

	// close response body stream at the end of the function
	defer resp.Body.Close()

	// read text contents of the HTTP response
	body, _ := io.ReadAll(resp.Body)

	// return number of bytes in response text
	return len(body)
}
