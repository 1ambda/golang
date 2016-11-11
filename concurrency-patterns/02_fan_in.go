package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're boring; I'm leaving")
}

func fanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)

	for i := range inputs {
		input := inputs[i]

		go func() {
			for {
				c <- <-input
			}
		}()
	}

	return c
}

func boring(msg string) chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
