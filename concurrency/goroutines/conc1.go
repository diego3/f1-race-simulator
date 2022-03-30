package goroutines

import (
	"fmt"
	"time"
)

// https://medium.com/@josemiguelmelo/goroutines-and-go-channels-ef04e8731f0e
func SimpleTest() {
	ch := make(chan int)

	// reader
	go func() {
		for {
			fmt.Printf("readingA %v\n", <-ch)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Printf("readingB %v\n", <-ch)
			time.Sleep(time.Second)
		}
	}()

	// writer
	for i := 0; i < 200; i++ {
		fmt.Printf("send request: %v\n", i)
		ch <- i
	}

	fmt.Println("waiting all requests to be processed")
	//time.Sleep(time.Second * 500)
}
