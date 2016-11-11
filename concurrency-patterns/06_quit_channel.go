package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	quit := make(chan bool)
	c := boring("Joe", quit)

	for i := rand.Intn(5); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- true
}

func boring(msg string, quit chan bool) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {

			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("I'm full")
				return
			}

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
