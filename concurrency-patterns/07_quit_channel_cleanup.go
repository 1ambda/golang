package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	quit := make(chan string)
	c := boring("Joe", quit)

	for i := rand.Intn(5); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "Finish it!"
	// do some cleanup
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {

			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				// do some cleanup
				quit <- "Thanks!"
				return
			}

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
