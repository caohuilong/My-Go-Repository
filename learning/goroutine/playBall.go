package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main2() {
	wg.Add(2)

	court := make(chan int)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
