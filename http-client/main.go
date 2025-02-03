package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := &http.Client{}
	requests := 1000
	url := "http://localhost:8080/"

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("Request error:", err)
				return
			}
			_, _ = ioutil.ReadAll(resp.Body)
			resp.Body.Close()
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("HTTP Throughput: %d requests in %v (%.2f req/sec)\n", requests, duration, float64(requests)/duration.Seconds())
}
