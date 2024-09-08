package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if resp != nil {
		_, _ = io.ReadAll(resp.Body)
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func main() {
	//url := "http://host.docker.internal:8080/user/1" // Change this to your local service URL
	url := "http://arch.homework/user/1"
	duration := 15 * time.Minute
	targetRPS := 100
	totalRequests := int(duration.Seconds()) * targetRPS

	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go sendRequest(url, &wg)
		// Control the rate
		if (i+1)%targetRPS == 0 {
			elapsed := time.Since(start)
			time.Sleep(time.Second - elapsed)
			start = time.Now()
		}
	}

	wg.Wait()
	fmt.Println("Test completed.")
}
