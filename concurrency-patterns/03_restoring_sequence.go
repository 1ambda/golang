package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg  string
	wait chan bool
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		msg2 := <-c
		fmt.Println(msg1.msg)
		fmt.Println(msg2.msg)
		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("You're boring; I'm leaving")
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)

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

func boring(msg string) <-chan Message {
	c := make(chan Message)
	wait := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-wait
		}
	}()

	return c
}
