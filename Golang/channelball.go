package main

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

func main() {

	court := make(chan int)

	wg.Add(2)
	go player("neo", court)
	go player("yang", court)

	court <- 1
	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Println(name, " is win")
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println(name, " is missed")
			close(court)
			return
		}
		fmt.Println(name, " hit ball ", ball)
		ball++
		court <- ball
	}
}
