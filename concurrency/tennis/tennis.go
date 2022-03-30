package tennis

// This sample program demonstrates how to use an unbuffered
// channel to simulate a game of tennis between two goroutines.

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

func Tennis() {
	wg.Add(2)
	court := make(chan int)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()
}

func player(name string, channel chan int) {
	defer wg.Done()

	for {
		time.Sleep(time.Second)

		ball, ok := <-channel
		if !ok {
			fmt.Printf("Player %s WON\n", name)

			return
		}

		r := rand.Intn(100)
		if r%13 == 0 {
			fmt.Printf("%s miss\n", name)
			close(channel)
			return
		}

		fmt.Printf("%s hit ball %d\n", name, ball)
		ball++

		channel <- ball
	}

}
